package FriendApp

import (
	SessionApp "github.com/slipe-fun/skid-backend/internal/app/session"
	FriendRepo "github.com/slipe-fun/skid-backend/internal/repository/friend"
	UserRepo "github.com/slipe-fun/skid-backend/internal/repository/user"
	"github.com/slipe-fun/skid-backend/internal/service"
)

type FriendApp struct {
	sessionApp *SessionApp.SessionApp
	friends    *FriendRepo.FriendRepo
	users      *UserRepo.UserRepo
	tokenSvc   *service.TokenService
}

func NewFriendApp(sessionApp *SessionApp.SessionApp,
	friends *FriendRepo.FriendRepo,
	users *UserRepo.UserRepo,
	tokenSvc *service.TokenService) *FriendApp {
	return &FriendApp{
		sessionApp: sessionApp,
		friends:    friends,
		users:      users,
		tokenSvc:   tokenSvc,
	}
}
