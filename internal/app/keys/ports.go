package keys

import (
	"github.com/slipe-fun/skid-backend/internal/domain"
)

type KeysRepo interface {
	GetByUserID(user_id int) (*domain.EncryptedKeys, error)
	Edit(keys *domain.EncryptedKeys) error
	Create(keys *domain.EncryptedKeys) (*domain.EncryptedKeys, error)
}
