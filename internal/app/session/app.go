package SessionApp

import (
	SessionRepo "github.com/slipe-fun/skid-backend/internal/repository/session"
	UserRepo "github.com/slipe-fun/skid-backend/internal/repository/user"
	"github.com/slipe-fun/skid-backend/internal/service"
)

type SessionApp struct {
	session  *SessionRepo.SessionRepo
	users    *UserRepo.UserRepo
	jwtSvc   *service.JWTService
	tokenSvc *service.TokenService
}

func NewSessionApp(session *SessionRepo.SessionRepo,
	users *UserRepo.UserRepo,
	jwtSvc *service.JWTService,
	tokenSvc *service.TokenService) *SessionApp {
	return &SessionApp{
		session:  session,
		users:    users,
		jwtSvc:   jwtSvc,
		tokenSvc: tokenSvc,
	}
}
