package friend

import "github.com/slipe-fun/skid-backend/internal/transport/ws/types"

type FriendHandler struct {
	friendApp FriendApp
	wsHub     *types.Hub
}

func NewFriendHandler(friendApp FriendApp, wsHub *types.Hub) *FriendHandler {
	return &FriendHandler{
		friendApp: friendApp,
		wsHub:     wsHub,
	}
}
