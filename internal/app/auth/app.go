package AuthApp

import (
	VerificationApp "github.com/slipe-fun/skid-backend/internal/app/verification"
	UserRepo "github.com/slipe-fun/skid-backend/internal/repository/user"
	VerificationRepo "github.com/slipe-fun/skid-backend/internal/repository/verification"
	"github.com/slipe-fun/skid-backend/internal/service"
	"github.com/slipe-fun/skid-backend/internal/service/oauth2"
)

type AuthApp struct {
	users     *UserRepo.UserRepo
	codesRepo *VerificationRepo.VerificationRepo
	codesApp  *VerificationApp.VerificationApp
	jwtSvc    *service.JWTService
	google    *oauth2.GoogleAuthService
}

func NewAuthApp(users *UserRepo.UserRepo,
	codesRepo *VerificationRepo.VerificationRepo,
	codesApp *VerificationApp.VerificationApp,
	jwt *service.JWTService,
	google *oauth2.GoogleAuthService) *AuthApp {
	return &AuthApp{
		users:     users,
		codesRepo: codesRepo,
		codesApp:  codesApp,
		jwtSvc:    jwt,
		google:    google,
	}
}
