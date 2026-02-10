package chat

import "github.com/slipe-fun/skid-backend/internal/domain"

func (c *ChatApp) HasMember(chat *domain.Chat, memberID int) bool {
	for _, m := range chat.Members {
		if m.ID == memberID {
			return true
		}
	}
	return false
}
