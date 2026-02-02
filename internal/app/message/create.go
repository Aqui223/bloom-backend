package MessageApp

import (
	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/service/logger"
)

func (m *MessageApp) CreateMessage(tokenStr string, chatId int, message *domain.Message) (*domain.Message, error) {
	_, err := m.sessionApp.GetSession(tokenStr)
	if err != nil {
		return nil, err
	}

	_, chatErr := m.chats.GetChatById(tokenStr, chatId)
	if chatErr != nil {
		return nil, chatErr
	}

	message, messageErr := m.messages.Create(message)
	if messageErr != nil {
		logger.LogError(messageErr.Error(), "message-app")
		return nil, domain.Failed("failed to create message")
	}

	return message, nil
}
