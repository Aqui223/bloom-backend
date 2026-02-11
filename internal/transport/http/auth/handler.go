package auth

type AuthHandler struct {
	authApp AuthApp
}

func NewAuthHandler(authApp AuthApp) *AuthHandler {
	return &AuthHandler{authApp: authApp}
}
