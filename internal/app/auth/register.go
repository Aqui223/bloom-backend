package AuthApp

import (
	"errors"

	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/service"
)

func (a *AuthApp) Register(username, password string) (string, *domain.User, error) {
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

	token, err := a.jwtSvc.GenerateToken(user.ID)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
