package friend

import "github.com/slipe-fun/skid-backend/internal/domain"

type FriendApp interface {
	SendRequest(string, int) (domain.FriendStatus, error)
	GetFriends(token, status string, limit, offset int) ([]domain.Friend, error)
	DeleteFriend(token string, friend_id int) error
}
