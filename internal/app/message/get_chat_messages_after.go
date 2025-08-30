package MessageApp

import (
	"github.com/slipe-fun/skid-backend/internal/domain"
)

func (c *MessageApp) GetChatMessagesAfter(tokenStr string, chatId int, afterId int) ([]*domain.Message, error) {
	_, err := c.tokenSvc.ExtractUserID(tokenStr)
	if err != nil {
		return nil, err
	}

	chat, err := c.chats.GetChatById(tokenStr, chatId)
	if err != nil {
		return nil, err
	}

	messages, err := c.messages.GetChatMessagesAfter(chat.ID, afterId)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
