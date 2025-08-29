package auth

import (
	"github.com/slipe-fun/skid-backend/internal/app"
)

type AuthHandler struct {
	authApp *app.AuthApp
}

func NewAuthHandler(authApp *app.AuthApp) *AuthHandler {
	return &AuthHandler{authApp: authApp}
}
