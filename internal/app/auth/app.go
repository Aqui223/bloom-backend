package AuthApp

import (
	UserRepo "github.com/slipe-fun/skid-backend/internal/repository/user"
	"github.com/slipe-fun/skid-backend/internal/service"
)

type AuthApp struct {
	users  *UserRepo.UserRepo
	jwtSvc *service.JWTService
}

func NewAuthApp(users *UserRepo.UserRepo, jwt *service.JWTService) *AuthApp {
	return &AuthApp{
		users:  users,
		jwtSvc: jwt,
	}
}
