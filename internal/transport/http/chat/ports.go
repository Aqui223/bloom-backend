package chat

import "github.com/slipe-fun/skid-backend/internal/domain"

type ChatApp interface {
	CreateChat(tokenStr string, recipient int) (*domain.Chat, *domain.Session, error)
	GetChatById(tokenStr string, id int) (*domain.Chat, error)
	GetChatsByUserId(tokenStr string) ([]*domain.ChatWithLastMessage, error)
	GetChatWithUsers(tokenStr string, recipient int) (*domain.Chat, error)
	GetOtherMember(chat *domain.Chat, memberID int) *domain.Member
	AddKeys(tokenStr string, chat *domain.Chat, kyberPublicKey string, ecdhPublicKey string, edPublicKey string) error
}

type MessageApp interface {
	GetChatMessages(tokenStr string, chatId int) ([]*domain.MessageWithReply, error)
	GetChatMessagesAfter(tokenStr string, chatId int, afterId int, count int) ([]*domain.MessageWithReply, error)
	GetChatMessagesBefore(tokenStr string, chatId int, afterId int, count int) ([]*domain.MessageWithReply, error)
	GetChatLastReadMessage(token string, chatId int) (*domain.Message, error)
}

type UserApp interface {
	GetUserById(id int) (*domain.User, error)
}
