package events

import (
	"encoding/json"

	"github.com/fasthttp/websocket"
	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/service/crypto"
	"github.com/slipe-fun/skid-backend/internal/transport/ws/types"
)

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
			Ciphertext:            message.Ciphertext,
			EncapsulatedKey:       message.EncapsulatedKey,
			Nonce:                 message.Nonce,
			ChatID:                message.ChatID,
			Signature:             message.Signature,
			SignedPayload:         message.SignedPayload,
			CEKWrap:               message.CEKWrap,
			CEKWrapIV:             message.CEKWrapIV,
			CEKWrapSalt:           message.CEKWrapSalt,
			EncapsulatedKeySender: message.EncapsulatedKeySender,
			CEKWrapSender:         message.CEKWrapSender,
			CEKWrapSenderIV:       message.CEKWrapSenderIV,
			CEKWrapSenderSalt:     message.CEKWrapSenderSalt,
		})
		if err != nil {
			SendError(sender, "failed_send_message")
			return
		}

		outMsg := struct {
			ID     int `json:"id"`
			UserID int `json:"user_id"`
			domain.SocketMessage
		}{
			ID:            sendedMessage.ID,
			UserID:        senderID,
			SocketMessage: message,
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
