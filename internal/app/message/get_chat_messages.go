package MessageApp

import (
	"github.com/slipe-fun/skid-backend/internal/domain"
)

func (c *MessageApp) GetChatMessages(tokenStr string, chatId int) ([]*domain.Message, error) {
	_, err := c.tokenSvc.ExtractUserID(tokenStr)
	if err != nil {
		return nil, err
	}

	chat, err := c.chats.GetChatById(tokenStr, chatId)
	if err != nil {
		return nil, err
	}

	messages, err := c.messages.GetChatMessages(chat.ID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
