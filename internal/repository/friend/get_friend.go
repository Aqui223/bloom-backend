package FriendRepo

import "github.com/slipe-fun/skid-backend/internal/domain"

func (r *FriendRepo) GetFriend(userID int, friendID int) (*domain.Friend, error) {
	query := `
		SELECT
			id,
			CASE
				WHEN user_id = $1 THEN friend_id
				ELSE user_id
			END AS friend_id,
			status
		FROM friends
		WHERE
			(user_id = $1 AND friend_id = $2)
			OR (user_id = $2 AND friend_id = $1)
		LIMIT 1
	`

	var friend domain.Friend
	err := r.db.Get(&friend, query, userID, friendID)
	if err != nil {
		return nil, err
	}

	return &friend, nil
}
