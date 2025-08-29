package auth

import (
	AuthApp "github.com/slipe-fun/skid-backend/internal/app/auth"
)

type AuthHandler struct {
	authApp *AuthApp.AuthApp
}

func NewAuthHandler(authApp *AuthApp.AuthApp) *AuthHandler {
	return &AuthHandler{authApp: authApp}
}
