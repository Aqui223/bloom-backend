package KeysApp

import (
	"encoding/base64"
	"errors"

	"github.com/slipe-fun/skid-backend/internal/domain"
)

func (k *KeysApp) CreateKeys(tokenStr string, chatId int, keys *domain.EncryptedKeys) (*domain.EncryptedKeys, error) {
	session, err := k.sessionApp.GetSession(tokenStr)
	if err != nil {
		return nil, errors.New("failed to get session")
	}

	chat, err := k.chats.GetChatById(tokenStr, chatId)
	if err != nil {
		return nil, errors.New("failed to get chat")
	}

	ciphertextBytes, err := base64.StdEncoding.DecodeString(keys.Ciphertext)
	if err != nil || len(ciphertextBytes) != 3393 {
		return nil, errors.New("invalid ciphertext")
	}

	nonceBytes, err := base64.StdEncoding.DecodeString(keys.Nonce)
	if err != nil || len(nonceBytes) != 12 {
		return nil, errors.New("invalid nonce")
	}

	saltBytes, err := base64.StdEncoding.DecodeString(keys.Salt)
	if err != nil || len(saltBytes) != 16 {
		return nil, errors.New("invalid salt")
	}

	keys, err = k.keys.Create(&domain.EncryptedKeys{
		UserID:     session.UserID,
		ChatID:     chat.ID,
		Ciphertext: keys.Ciphertext,
		Nonce:      keys.Nonce,
		Salt:       keys.Salt,
	})
	if err != nil {
		return nil, errors.New("failed to save keys")
	}

	return keys, nil
}
