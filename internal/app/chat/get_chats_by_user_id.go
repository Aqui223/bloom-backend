package ChatApp

import "github.com/slipe-fun/skid-backend/internal/domain"

func (c *ChatApp) GetChatsByUserId(tokenStr string) ([]*domain.Chat, error) {
	userID, err := c.tokenSvc.ExtractUserID(tokenStr)
	if err != nil {
		return nil, err
	}

	chats, err := c.chats.GetByUserId(userID)

	if err != nil {
		return nil, err
	}

	return chats, nil
}
