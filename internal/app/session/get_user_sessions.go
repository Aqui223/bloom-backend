package SessionApp

import (
	"errors"

	"github.com/slipe-fun/skid-backend/internal/domain"
)

func (s *SessionApp) GetUserSessions(user_id int) ([]*domain.Session, error) {
	user, err := s.users.GetById(user_id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	sessions, err := s.session.GetByUserId(user.ID)
	if err != nil {
		return nil, errors.New("failed to get user sessions")
	}

	return sessions, nil
}
