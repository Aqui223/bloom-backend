package user

type UserHandler struct {
	userApp   UserApp
	friendApp FriendApp
}

func NewUserHandler(userApp UserApp,
	friendApp FriendApp) *UserHandler {
	return &UserHandler{
		userApp:   userApp,
		friendApp: friendApp,
	}
}
