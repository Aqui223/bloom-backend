package chat

import (
	ChatApp "github.com/slipe-fun/skid-backend/internal/app/chat"
	MessageApp "github.com/slipe-fun/skid-backend/internal/app/message"
	UserApp "github.com/slipe-fun/skid-backend/internal/app/user"
)

type ChatHandler struct {
	chatApp    *ChatApp.ChatApp
	userApp    *UserApp.UserApp
	messageApp *MessageApp.MessageApp
}

func NewChatHandler(chatApp *ChatApp.ChatApp, userApp *UserApp.UserApp, messageApp *MessageApp.MessageApp) *ChatHandler {
	return &ChatHandler{
		chatApp:    chatApp,
		userApp:    userApp,
		messageApp: messageApp,
	}
}
