package keys

type KeysHandler struct {
	keysApp KeysApp
}

func NewKeysHandler(keysApp KeysApp) *KeysHandler {
	return &KeysHandler{
		keysApp: keysApp,
	}
}
