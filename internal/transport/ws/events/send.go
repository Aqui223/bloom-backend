package events

import (
	"encoding/json"
	"fmt"

	"github.com/fasthttp/websocket"
	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/service/crypto"
	"github.com/slipe-fun/skid-backend/internal/transport/ws/types"
)

func strPtrOrNil(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func intPtrOrNil(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

func Send(hub *types.Hub, sender *types.Client, token string, senderID int, room string, message domain.SocketMessage) {
	if clients, ok := hub.Clients[room]; ok {
		chat, err := hub.Chats.GetChatById(token, message.ChatID)
		if err != nil || chat == nil {
			SendError(sender, "chat_not_found")
			return
		}

		var member *domain.Member
		for i, m := range chat.Members {
			if m.ID == senderID {
				member = &chat.Members[i]
				break
			}
		}
		if member == nil {
			SendError(sender, "not_member")
			return
		}

		var replyTo *domain.Message
		if message.ReplyTo != 0 {
			replyToMessage, err := hub.Messages.GetMessageById(token, message.ReplyTo)
			if err != nil || replyToMessage == nil || replyToMessage.ChatID != chat.ID {
				SendError(sender, "reply_to_not_found")
				return
			}
			replyTo = replyToMessage
		}

		switch message.EncryptionType {
		case "server":
			sendedMessage, err := hub.Messages.CreateMessage(token, message.ChatID, &domain.Message{
				Ciphertext: message.Ciphertext,
				Nonce:      message.Nonce,
				ChatID:     message.ChatID,
				ReplyTo:    intPtrOrNil(message.ReplyTo),
			})

			if err != nil {
				fmt.Println(err)
				SendError(sender, "failed_send_message")
				return
			}

			outMsg := struct {
				Type           string          `json:"type"`
				EncryptionType string          `json:"encryption_type"`
				ID             int             `json:"id"`
				UserID         int             `json:"user_id"`
				ReplyTo        *domain.Message `json:"reply_to,omitempty"`
				domain.SocketMessage
			}{
				Type:           "message",
				EncryptionType: "server",
				ID:             sendedMessage.ID,
				UserID:         senderID,
				ReplyTo:        replyTo,
				SocketMessage:  message,
			}

			b, err := json.Marshal(outMsg)
			if err != nil {
				return
			}

			for client := range clients {
				if err := client.Conn.WriteMessage(websocket.TextMessage, b); err != nil {
					SendError(sender, "failed_send_message")
				}
			}
		case "client":
			if err := crypto.VerifySignature(
				member.EdPublicKey,
				message.SignedPayload,
				message.Signature,
			); err != nil {
				SendError(sender, "failed_verify_signature")
				return
			}

			if err := crypto.ValidateCEKFields(
				message.CEKWrap,
				message.CEKWrapIV,
				message.CEKWrapSalt,
				message.EncapsulatedKeySender,
				message.CEKWrapSender,
				message.CEKWrapSenderIV,
				message.CEKWrapSenderSalt,
			); err != nil {
				return
			}

			sendedMessage, err := hub.Messages.CreateMessage(token, message.ChatID, &domain.Message{
				Ciphertext: message.Ciphertext,
				Nonce:      message.Nonce,
				ChatID:     message.ChatID,
				ReplyTo:    intPtrOrNil(message.ReplyTo),

				EncapsulatedKey:       strPtrOrNil(message.EncapsulatedKey),
				Signature:             strPtrOrNil(message.Signature),
				SignedPayload:         strPtrOrNil(message.SignedPayload),
				CEKWrap:               strPtrOrNil(message.CEKWrap),
				CEKWrapIV:             strPtrOrNil(message.CEKWrapIV),
				CEKWrapSalt:           strPtrOrNil(message.CEKWrapSalt),
				EncapsulatedKeySender: strPtrOrNil(message.EncapsulatedKeySender),
				CEKWrapSender:         strPtrOrNil(message.CEKWrapSender),
				CEKWrapSenderIV:       strPtrOrNil(message.CEKWrapSenderIV),
				CEKWrapSenderSalt:     strPtrOrNil(message.CEKWrapSenderSalt),
			})

			if err != nil {
				SendError(sender, "failed_send_message")
				return
			}

			outMsg := struct {
				Type           string          `json:"type"`
				EncryptionType string          `json:"encryption_type"`
				ID             int             `json:"id"`
				UserID         int             `json:"user_id"`
				ReplyTo        *domain.Message `json:"reply_to,omitempty"`
				domain.SocketMessage
			}{
				Type:           "message",
				EncryptionType: "client",
				ID:             sendedMessage.ID,
				UserID:         senderID,
				ReplyTo:        replyTo,
				SocketMessage:  message,
			}

			b, err := json.Marshal(outMsg)
			if err != nil {
				return
			}

			for client := range clients {
				if err := client.Conn.WriteMessage(websocket.TextMessage, b); err != nil {
					SendError(sender, "failed_send_message")
				}
			}
		}
	}
}
