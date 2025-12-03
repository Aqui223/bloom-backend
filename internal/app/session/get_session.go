package SessionApp

import (
	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/service"
	"github.com/slipe-fun/skid-backend/internal/service/logger"
)

func (s *SessionApp) GetSession(token string) (*domain.Session, error) {
	userID, err := s.tokenSvc.ExtractUserID(token)
	if err != nil {
		logger.LogError(err.Error(), "session-app")
		return nil, domain.Unauthorized("failed to extract token")
	}

	_, err = s.users.GetById(userID)
	if err != nil {
		logger.LogError(err.Error(), "session-app")
		return nil, domain.Unauthorized("user not found")
	}

	session, err := s.session.GetByToken(service.HashSHA256(token))
	if err != nil {
		logger.LogError(err.Error(), "session-app")
		return nil, domain.Unauthorized("session not found")
	}

	if session.UserID != userID {
		return nil, domain.Unauthorized("session not found")
	}

	return session, nil
}
