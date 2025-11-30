package keys

import (
	ChatApp "github.com/slipe-fun/skid-backend/internal/app/chat"
	KeysApp "github.com/slipe-fun/skid-backend/internal/app/keys"
)

type KeysHandler struct {
	keysApp *KeysApp.KeysApp
	chatApp *ChatApp.ChatApp
}

func NewKeysHandler(keysApp *KeysApp.KeysApp, chatApp *ChatApp.ChatApp) *KeysHandler {
	return &KeysHandler{
		keysApp: keysApp,
		chatApp: chatApp,
	}
}
