package UserApp

import "github.com/slipe-fun/skid-backend/internal/domain"

func (u *UserApp) SearchUsersByUsername(username string, limit, offset int) ([]*domain.User, error) {
	return u.users.SearchUsersByUsername(username, limit, offset)
}
