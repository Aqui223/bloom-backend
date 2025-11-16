package SessionApp

import (
	"errors"

	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/service"
)

func (s *SessionApp) GetSession(token string) (*domain.Session, error) {
	userID, err := s.tokenSvc.ExtractUserID(token)
	if err != nil {
		return nil, err
	}

	_, err = s.users.GetById(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	session, err := s.session.GetByToken(service.HashSHA256(token))
	if err != nil {
		return nil, errors.New("session not found")
	}

	if session.UserID != userID {
		return nil, errors.New("session not found")
	}

	return session, nil
}
