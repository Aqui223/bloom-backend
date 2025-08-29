package MessageApp

import (
	"github.com/slipe-fun/skid-backend/internal/domain"
)

func (c *MessageApp) GetMessageById(tokenStr string, id int) (*domain.Message, error) {
	_, err := c.tokenSvc.ExtractUserID(tokenStr)
	if err != nil {
		return nil, err
	}

	message, err := c.messages.GetById(id)
	if err != nil {
		return nil, err
	}

	_, chaterr := c.chats.GetChatById(tokenStr, message.ChatID)
	if chaterr != nil {
		return nil, chaterr
	}

	return message, nil
}
