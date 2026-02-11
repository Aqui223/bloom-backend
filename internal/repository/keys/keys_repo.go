package keys

import (
	"github.com/jmoiron/sqlx"
)

type KeysRepo struct {
	db *sqlx.DB
}

func NewKeysRepo(db *sqlx.DB) *KeysRepo {
	return &KeysRepo{db: db}
}
