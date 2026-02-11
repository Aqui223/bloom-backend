package session

import (
	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/pkg/logger"
)

func (s *SessionApp) GetUserSessions(user_id int) ([]*domain.Session, error) {
	user, err := s.users.GetByID(user_id)
	if err != nil {
		logger.LogError(err.Error(), "session-app")
		return nil, domain.NotFound("user not found")
	}

	sessions, err := s.session.GetByUserID(user.ID)
	if err != nil {
		logger.LogError(err.Error(), "session-app")
		return nil, domain.Failed("failed to get user sessions")
	}

	return sessions, nil
}
