package friend

import (
	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/pkg/logger"
)

func (f *FriendApp) GetFriends(user_id int, status string, limit, offset int) ([]domain.Friend, error) {
	friends, err := f.friends.GetFriends(user_id, status, limit, offset)
	if err != nil {
		logger.LogError(err.Error(), "friend-app")
		return nil, domain.Failed("failed to get friends")
	}

	return friends, nil
}
