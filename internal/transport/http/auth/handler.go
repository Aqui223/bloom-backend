package auth

import (
	AuthApp "github.com/slipe-fun/skid-backend/internal/app/auth"
	"github.com/slipe-fun/skid-backend/internal/service/oauth2"
)

type AuthHandler struct {
	authApp *AuthApp.AuthApp
	google  *oauth2.GoogleAuthService
}

func NewAuthHandler(authApp *AuthApp.AuthApp, google *oauth2.GoogleAuthService) *AuthHandler {
	return &AuthHandler{authApp: authApp, google: google}
}
