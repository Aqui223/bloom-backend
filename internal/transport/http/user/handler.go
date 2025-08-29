package user

import (
	"github.com/slipe-fun/skid-backend/internal/app"
)

type UserHandler struct {
	userApp *app.UserApp
}

func NewUserHandler(userApp *app.UserApp) *UserHandler {
	return &UserHandler{userApp: userApp}
}
