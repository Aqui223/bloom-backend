package friend

import (
	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/pkg/logger"
)

func (f *FriendApp) DeleteFriend(user_id, friend_id int) error {
	_, err := f.friends.GetFriend(user_id, friend_id)
	if err != nil {
		logger.LogError(err.Error(), "friend-app")
		return domain.NotFound("friend not found")
	}

	err = f.friends.Delete(user_id, friend_id)
	if err != nil {
		logger.LogError(err.Error(), "friend-app")
		return domain.Failed("failed to delete friend")
	}

	return nil
}
