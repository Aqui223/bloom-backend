package friend

import FriendApp "github.com/slipe-fun/skid-backend/internal/app/friend"

type FriendHandler struct {
	friendApp *FriendApp.FriendApp
}

func NewFriendHandler(friendApp *FriendApp.FriendApp) *FriendHandler {
	return &FriendHandler{
		friendApp: friendApp,
	}
}
