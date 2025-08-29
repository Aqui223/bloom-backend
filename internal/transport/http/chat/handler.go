package chat

import (
	"github.com/slipe-fun/skid-backend/internal/app"
)

type ChatHandler struct {
	chatApp *app.ChatApp
	userApp *app.UserApp
}

func NewChatHandler(chatApp *app.ChatApp, userApp *app.UserApp) *ChatHandler {
	return &ChatHandler{
		chatApp: chatApp,
		userApp: userApp,
	}
}
