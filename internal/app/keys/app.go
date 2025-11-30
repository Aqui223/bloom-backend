package KeysApp

import (
	ChatApp "github.com/slipe-fun/skid-backend/internal/app/chat"
	SessionApp "github.com/slipe-fun/skid-backend/internal/app/session"
	UserApp "github.com/slipe-fun/skid-backend/internal/app/user"
	KeysRepo "github.com/slipe-fun/skid-backend/internal/repository/keys"
)

type KeysApp struct {
	sessionApp *SessionApp.SessionApp
	keys       *KeysRepo.KeysRepo
	users      *UserApp.UserApp
	chats      *ChatApp.ChatApp
}

func NewKeysApp(sessionApp *SessionApp.SessionApp,
	keys *KeysRepo.KeysRepo,
	users *UserApp.UserApp,
	chats *ChatApp.ChatApp) *KeysApp {
	return &KeysApp{
		sessionApp: sessionApp,
		keys:       keys,
		users:      users,
		chats:      chats,
	}
}
