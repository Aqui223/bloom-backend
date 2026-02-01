package user

import (
	FriendApp "github.com/slipe-fun/skid-backend/internal/app/friend"
	UserApp "github.com/slipe-fun/skid-backend/internal/app/user"
)

type UserHandler struct {
	userApp   *UserApp.UserApp
	friendApp *FriendApp.FriendApp
}

func NewUserHandler(userApp *UserApp.UserApp,
	friendApp *FriendApp.FriendApp) *UserHandler {
	return &UserHandler{
		userApp:   userApp,
		friendApp: friendApp,
	}
}
