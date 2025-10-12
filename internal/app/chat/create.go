package ChatApp

import (
	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/service/crypto"
)

func (c *ChatApp) CreateChat(tokenStr string, recipient int) (*domain.Chat, error) {
	userID, err := c.tokenSvc.ExtractUserID(tokenStr)
	if err != nil {
		return nil, err
	}

	encryptionKey, err := crypto.GenerateEncryptionKey()
	if err != nil {
		return nil, err
	}

	encKey := encryptionKey
	chat, err := c.chats.Create(&domain.Chat{
		Members: []domain.Member{
			{
				ID:             userID,
				KyberPublicKey: "",
				EcdhPublicKey:  "",
				EdPublicKey:    "",
			},
			{
				ID:             recipient,
				KyberPublicKey: "",
				EcdhPublicKey:  "",
				EdPublicKey:    "",
			},
		},
		EncryptionKey: &encKey,
	})

	if err != nil {
		return nil, err
	}

	return chat, nil
}
