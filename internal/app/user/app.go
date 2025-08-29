package UserApp

import (
	UserRepo "github.com/slipe-fun/skid-backend/internal/repository/user"
	"github.com/slipe-fun/skid-backend/internal/service"
)

type UserApp struct {
	users    *UserRepo.UserRepo
	jwtSvc   *service.JWTService
	tokenSvc *service.TokenService
}

func NewUserApp(users *UserRepo.UserRepo, jwtSvc *service.JWTService, tokenSvc *service.TokenService) *UserApp {
	return &UserApp{
		users:    users,
		jwtSvc:   jwtSvc,
		tokenSvc: tokenSvc,
	}
}
