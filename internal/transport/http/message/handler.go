package message

import (
	ChatApp "github.com/slipe-fun/skid-backend/internal/app/chat"
	MessageApp "github.com/slipe-fun/skid-backend/internal/app/message"
	UserApp "github.com/slipe-fun/skid-backend/internal/app/user"
)

type MessageHandler struct {
	chatApp    *ChatApp.ChatApp
	userApp    *UserApp.UserApp
	messageApp *MessageApp.MessageApp
}

func NewMessageHandler(chatApp *ChatApp.ChatApp, userApp *UserApp.UserApp, messageApp *MessageApp.MessageApp) *MessageHandler {
	return &MessageHandler{
		chatApp:    chatApp,
		userApp:    userApp,
		messageApp: messageApp,
	}
}
