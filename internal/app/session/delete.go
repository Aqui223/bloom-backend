package SessionApp

import "errors"

func (s *SessionApp) DeleteSession(id int, token string) error {
	userID, err := s.tokenSvc.ExtractUserID(token)
	if err != nil {
		return err
	}

	_, err = s.users.GetById(userID)
	if err != nil {
		return errors.New("user not found")
	}

	session, err := s.session.GetById(id)
	if err != nil {
		return errors.New("session not found")
	}

	if session.UserID != userID {
		return errors.New("session not found")
	}

	deleteSessionErr := s.session.Delete(id)
	if deleteSessionErr != nil {
		return errors.New("failed to delete session")
	}

	return nil
}
