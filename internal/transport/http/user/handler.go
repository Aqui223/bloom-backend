package user

import (
	UserApp "github.com/slipe-fun/skid-backend/internal/app/user"
)

type UserHandler struct {
	userApp *UserApp.UserApp
}

func NewUserHandler(userApp *UserApp.UserApp) *UserHandler {
	return &UserHandler{userApp: userApp}
}
