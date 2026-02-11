package friend

type FriendHandler struct {
	friendApp FriendApp
}

func NewFriendHandler(friendApp FriendApp) *FriendHandler {
	return &FriendHandler{
		friendApp: friendApp,
	}
}
