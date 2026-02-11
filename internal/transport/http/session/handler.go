package session

type SessionHandler struct {
	sessionApp SessionApp
}

func NewSessionHandler(sessionApp SessionApp) *SessionHandler {
	return &SessionHandler{
		sessionApp: sessionApp,
	}
}
