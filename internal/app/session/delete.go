package session

import (
	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/pkg/logger"
)

func (s *SessionApp) DeleteSession(user_id, id int) error {
	session, err := s.session.GetByID(id)
	if err != nil {
		logger.LogError(err.Error(), "session-app")
		return domain.NotFound("session not found")
	}

	if session.UserID != user_id {
		return domain.NotFound("session not found")
	}

	deleteSessionErr := s.session.Delete(id)
	if deleteSessionErr != nil {
		logger.LogError(deleteSessionErr.Error(), "session-app")
		return domain.Failed("failed to delete session")
	}

	return nil
}
