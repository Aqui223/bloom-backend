package ChatApp

import (
	"errors"

	"github.com/slipe-fun/skid-backend/internal/domain"
)

func (c *ChatApp) GetChatById(tokenStr string, id int) (*domain.Chat, error) {
	userID, err := c.tokenSvc.ExtractUserID(tokenStr)

	if err != nil {
		return nil, err
	}

	chat, err := c.chats.GetById(id)

	if err != nil {
		return nil, err
	}

	if !c.HasMember(chat, userID) {
		return nil, errors.New("user is not in chat")
	}

	return chat, nil
}
