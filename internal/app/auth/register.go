package AuthApp

import (
	"errors"
	"time"

	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/service"
)

func (a *AuthApp) Register(username, password string, expire time.Duration) (string, *domain.User, error) {
	_, err := a.users.GetByUsername(username)

	if err == nil {
		return "", nil, errors.New("user already exists")
	}

	hashedPassword, err := service.HashPassword(password)

	if err != nil {
		return "", nil, err
	}

	user, err := a.users.Create(&domain.User{
		Username: username,
		Password: hashedPassword,
	})

	if err != nil {
		return "", nil, err
	}

	token, err := a.jwtSvc.GenerateToken(user.ID, expire)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
