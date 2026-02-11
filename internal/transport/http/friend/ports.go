package friend

import "github.com/slipe-fun/skid-backend/internal/domain"

type FriendApp interface {
	SendRequest(user_id, receiverID int) (domain.FriendStatus, error)
	GetFriends(user_id int, status string, limit, offset int) ([]domain.Friend, error)
	DeleteFriend(user_id, friend_id int) error
}
