package auth

import "github.com/slipe-fun/skid-backend/internal/domain"

type AuthApp interface {
	RequestCode(email string) error
	ExchangeCode(code string) (string, *domain.User, error)
	VerifyCode(email string, code string) (string, *domain.User, error)
	Register(email string) error
}
