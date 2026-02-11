package chat

import (
	"github.com/jmoiron/sqlx"
)

type ChatRepo struct {
	db       *sqlx.DB
	userRepo UserRepo
}

func NewChatRepo(db *sqlx.DB, userRepo UserRepo) *ChatRepo {
	return &ChatRepo{db: db, userRepo: userRepo}
}
