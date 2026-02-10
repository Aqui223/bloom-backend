package message

import (
	ChatApp "github.com/slipe-fun/skid-backend/internal/app/chat"
	SessionApp "github.com/slipe-fun/skid-backend/internal/app/session"
	MessageRepo "github.com/slipe-fun/skid-backend/internal/repository/message"
	"github.com/slipe-fun/skid-backend/internal/service"
)

type MessageApp struct {
	sessionApp *SessionApp.SessionApp
	messages   *MessageRepo.MessageRepo
	chats      *ChatApp.ChatApp
	tokenSvc   *service.TokenService
}

func NewMessageApp(sessionApp *SessionApp.SessionApp,
	messages *MessageRepo.MessageRepo,
	chats *ChatApp.ChatApp,
	tokenSvc *service.TokenService) *MessageApp {
	return &MessageApp{
		sessionApp: sessionApp,
		messages:   messages,
		chats:      chats,
		tokenSvc:   tokenSvc,
	}
}
