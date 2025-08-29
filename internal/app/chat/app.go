package ChatApp

import (
	ChatRepo "github.com/slipe-fun/skid-backend/internal/repository/chat"
	"github.com/slipe-fun/skid-backend/internal/service"
)

type ChatApp struct {
	chats    *ChatRepo.ChatRepo
	tokenSvc *service.TokenService
}

func NewChatApp(chats *ChatRepo.ChatRepo, tokenSvc *service.TokenService) *ChatApp {
	return &ChatApp{
		chats:    chats,
		tokenSvc: tokenSvc,
	}
}
