package ChatApp

import "github.com/slipe-fun/skid-backend/internal/domain"

func (c *ChatApp) AddKeys(tokenStr string, chat *domain.Chat, kyberPublicKey string, ecdhPublicKey string, edPublicKey string) error {
	userID, err := c.tokenSvc.ExtractUserID(tokenStr)
	if err != nil {
		return err
	}

	for i := range chat.Members {
		if chat.Members[i].ID == userID {
			chat.Members[i].KyberPublicKey = kyberPublicKey
			chat.Members[i].EcdhPublicKey = ecdhPublicKey
			chat.Members[i].EdPublicKey = edPublicKey
		}
	}

	return c.chats.UpdateChat(chat)
}
