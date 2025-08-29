package UserApp

import "github.com/slipe-fun/skid-backend/internal/domain"

func (u *UserApp) GetUserByToken(tokenStr string) (*domain.User, error) {
	userID, err := u.tokenSvc.ExtractUserID(tokenStr)
	if err != nil {
		return nil, err
	}

	user, err := u.users.GetById(int(userID))
	if err != nil {
		return nil, err
	}

	return user, nil
}
