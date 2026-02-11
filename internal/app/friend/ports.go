package friend

import "github.com/slipe-fun/skid-backend/internal/domain"

type FriendRepo interface {
	GetFriend(userID, friendID int) (*domain.FriendRow, error)
	Delete(userID, friendID int) error
	GetFriendCount(userID int) (int, error)
	GetFriends(userID int, status string, limit, offset int) ([]domain.Friend, error)
	EditStatus(userID, friendID int, status domain.FriendStatus) error
	Create(friend *domain.FriendRow) (*domain.FriendRow, error)
}

type UserRepo interface {
	GetByID(id int) (*domain.User, error)
}
