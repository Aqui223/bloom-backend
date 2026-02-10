package session

import (
	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/service"
	"github.com/slipe-fun/skid-backend/internal/service/logger"
)

func (s *SessionApp) GetUserSessions(token string) ([]*domain.Session, error) {
	session, err := s.session.GetByToken(service.HashSHA256(token))
	if err != nil {
		logger.LogError(err.Error(), "session-app")
		return nil, domain.Unauthorized("session not found")
	}

	user, err := s.users.GetById(session.UserID)
	if err != nil {
		logger.LogError(err.Error(), "session-app")
		return nil, domain.NotFound("user not found")
	}

	sessions, err := s.session.GetByUserId(user.ID)
	if err != nil {
		logger.LogError(err.Error(), "session-app")
		return nil, domain.Failed("failed to get user sessions")
	}

	return sessions, nil
}
