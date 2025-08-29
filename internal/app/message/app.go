package MessageApp

import (
	ChatApp "github.com/slipe-fun/skid-backend/internal/app/chat"
	MessageRepo "github.com/slipe-fun/skid-backend/internal/repository/message"
	"github.com/slipe-fun/skid-backend/internal/service"
)

type MessageApp struct {
	messages *MessageRepo.MessageRepo
	chats    *ChatApp.ChatApp
	tokenSvc *service.TokenService
}

func NewMessageApp(messages *MessageRepo.MessageRepo, chats *ChatApp.ChatApp, tokenSvc *service.TokenService) *MessageApp {
	return &MessageApp{
		messages: messages,
		chats:    chats,
		tokenSvc: tokenSvc,
	}
}
