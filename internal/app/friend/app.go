package friend

type FriendApp struct {
	friends FriendRepo
	users   UserRepo
}

func NewFriendApp(friends FriendRepo,
	users UserRepo,
) *FriendApp {
	return &FriendApp{
		friends: friends,
		users:   users,
	}
}
