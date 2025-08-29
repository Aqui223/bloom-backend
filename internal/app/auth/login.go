package AuthApp

import (
	"errors"

	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/service"
)

func (a *AuthApp) Login(username, password string) (string, *domain.User, error) {
	user, err := a.users.GetByUsername(username)

	if err != nil {
		return "", nil, errors.New("user not found")
	}

	ok, err := service.VerifyPassword(password, user.Password)

	if err != nil || !ok {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := a.jwtSvc.GenerateToken(user.ID)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
