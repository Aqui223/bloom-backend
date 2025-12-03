package MessageApp

import (
	"time"

	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/service/logger"
)

func (m *MessageApp) UpdateMessagesSeenStatus(messageIDs []int, seenAt time.Time) error {
	updateMessageError := m.messages.UpdateMessagesSeenStatus(messageIDs, seenAt)
	if updateMessageError != nil {
		logger.LogError(updateMessageError.Error(), "message-app")
		return domain.Failed("failed to update message seen status")
	}

	return nil
}
