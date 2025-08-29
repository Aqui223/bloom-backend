package ChatApp

import "github.com/slipe-fun/skid-backend/internal/domain"

func (c *ChatApp) GetChatWithUsers(tokenStr string, recipient int) (*domain.Chat, error) {
	userID, err := c.tokenSvc.ExtractUserID(tokenStr)
	if err != nil {
		return nil, err
	}

	chat, err := c.chats.GetWithUsers(userID, recipient)

	if err != nil {
		return nil, err
	}

	return chat, nil
}
