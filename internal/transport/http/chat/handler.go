package chat

import "github.com/slipe-fun/skid-backend/internal/transport/ws/types"

type ChatHandler struct {
	chatApp    ChatApp
	userApp    UserApp
	messageApp MessageApp
	wsHub      *types.Hub
}

func NewChatHandler(chatApp ChatApp, userApp UserApp, messageApp MessageApp, wsHub *types.Hub) *ChatHandler {
	return &ChatHandler{
		chatApp:    chatApp,
		userApp:    userApp,
		messageApp: messageApp,
		wsHub:      wsHub,
	}
}
