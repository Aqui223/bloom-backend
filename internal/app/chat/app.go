package ChatApp

import (
	SessionApp "github.com/slipe-fun/skid-backend/internal/app/session"
	ChatRepo "github.com/slipe-fun/skid-backend/internal/repository/chat"
	"github.com/slipe-fun/skid-backend/internal/service"
)

type ChatApp struct {
	sessionApp *SessionApp.SessionApp
	chats      *ChatRepo.ChatRepo
	tokenSvc   *service.TokenService
}

func NewChatApp(sessionApp *SessionApp.SessionApp, chats *ChatRepo.ChatRepo, tokenSvc *service.TokenService) *ChatApp {
	return &ChatApp{
		sessionApp: sessionApp,
		chats:      chats,
		tokenSvc:   tokenSvc,
	}
}
