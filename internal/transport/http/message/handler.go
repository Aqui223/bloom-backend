package message

import (
	"github.com/slipe-fun/skid-backend/internal/transport/ws/types"
)

type MessageHandler struct {
	chatApp    ChatApp
	messageApp MessageApp
	wsHub      *types.Hub
}

func NewMessageHandler(chatApp ChatApp, messageApp MessageApp, wsHub *types.Hub) *MessageHandler {
	return &MessageHandler{
		chatApp:    chatApp,
		messageApp: messageApp,
		wsHub:      wsHub,
	}
}
