package session

import "github.com/slipe-fun/skid-backend/internal/domain"

type SessionApp interface {
	DeleteSession(id int, token string) error
	GetSession(token string) (*domain.Session, error)
	GetUserSessions(token string) ([]*domain.Session, error)
}
