package events

import (
	"encoding/json"
	"log"

	"github.com/fasthttp/websocket"
	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/transport/ws/types"
)

func Send(hub *types.Hub, sender *types.Client, token string, senderID int, room string, message domain.SocketMessage) {
	if clients, ok := hub.Clients[room]; ok {
		for client := range clients {
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
				log.Println("Failed to send message:", err)
				continue
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
				log.Println("Failed to marshal message:", err)
				continue
			}

			if err := client.Conn.WriteMessage(websocket.TextMessage, b); err != nil {
				log.Println("Failed to send message:", err)
			}
		}
	}
}
