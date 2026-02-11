package keys

import "github.com/slipe-fun/skid-backend/internal/domain"

type KeysApp interface {
	CreateKeys(tokenStr string, keys *domain.EncryptedKeys) (*domain.EncryptedKeys, error)
	GetUserChatsKeys(tokenStr string) (*domain.EncryptedKeys, error)
}
