package friend

import (
	"database/sql"
	"errors"

	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/pkg/logger"
)

func (f *FriendApp) SendRequest(token string, receiverID int) (domain.FriendStatus, error) {
	session, err := f.sessionApp.GetSession(token)
	if err != nil {
		return domain.FriendStatus(""), err
	}

	if session.UserID == receiverID {
		return domain.FriendStatus(""), domain.InvalidData("cannot add yourself as friend")
	}

	receiver, err := f.users.GetByID(receiverID)
	if err != nil {
		logger.LogError(err.Error(), "user-app")
		return domain.FriendStatus(""), domain.NotFound("receiver not found")
	}

	friend, err := f.friends.GetFriend(session.UserID, receiver.ID)
	if err == nil {
		if friend.Status == "pending" {
			if friend.FriendID == session.UserID {
				if err := f.friends.EditStatus(session.UserID, receiverID, domain.StatusAccepted); err != nil {
					logger.LogError(err.Error(), "friend-app")
					return domain.FriendStatus(""), domain.Failed("failed to accept friend request")
				}
				return domain.StatusAccepted, nil
			}

			return domain.FriendStatus(""), domain.AlreadyExists("friend request already sent")
		}

		if friend.Status == domain.StatusAccepted {
			return domain.FriendStatus(""), domain.AlreadyExists("already friends")
		}

		return domain.FriendStatus(""), domain.AlreadyExists("invalid friend state")
	}

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.LogError(err.Error(), "friend-app")
		return domain.FriendStatus(""), domain.Failed("failed to check friend relation")
	}

	_, err = f.friends.Create(&domain.FriendRow{
		UserID:   session.UserID,
		FriendID: receiver.ID,
		Status:   domain.StatusPending,
	})
	if err != nil {
		logger.LogError(err.Error(), "user-app")
		return domain.FriendStatus(""), domain.Failed("failed to send friend request")
	}

	return domain.StatusPending, nil
}
