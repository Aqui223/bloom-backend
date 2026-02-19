package encryptedchatkeys

import "github.com/slipe-fun/skid-backend/internal/domain"

func (k *EncryptedChatKeysApp) GetBySessionID(session_id int) ([]*domain.EncryptedChatKeys, error) {
	keys, err := k.keys.GetBySessionID(session_id)
	if err != nil {
		return nil, domain.Failed("failed to get sessions")
	}

	return keys, nil
}
