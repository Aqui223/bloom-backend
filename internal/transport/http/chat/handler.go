package chat

import (
	ChatApp "github.com/slipe-fun/skid-backend/internal/app/chat"
	UserApp "github.com/slipe-fun/skid-backend/internal/app/user"
)

type ChatHandler struct {
	chatApp *ChatApp.ChatApp
	userApp *UserApp.UserApp
}

func NewChatHandler(chatApp *ChatApp.ChatApp, userApp *UserApp.UserApp) *ChatHandler {
	return &ChatHandler{
		chatApp: chatApp,
		userApp: userApp,
	}
}
