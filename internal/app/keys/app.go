package keys

type KeysApp struct {
	keys KeysRepo
}

func NewKeysApp(keys KeysRepo) *KeysApp {
	return &KeysApp{
		keys: keys,
	}
}
