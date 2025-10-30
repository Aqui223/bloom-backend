package domain

type SocketMessage struct {
	ChatID                int    `json:"chat_id"`
	EncryptionType        string `json:"encryption_type"`
	Ciphertext            string `json:"ciphertext"`
	Nonce                 string `json:"nonce"`
	EncapsulatedKey       string `json:"encapsulated_key,omitempty"`
	Signature             string `json:"signature,omitempty"`
	SignedPayload         string `json:"signed_payload,omitempty"`
	CEKWrap               string `json:"cek_wrap,omitempty"`
	CEKWrapIV             string `json:"cek_wrap_iv,omitempty"`
	CEKWrapSalt           string `json:"cek_wrap_salt,omitempty"`
	EncapsulatedKeySender string `json:"encapsulated_key_sender,omitempty"`
	CEKWrapSender         string `json:"cek_wrap_sender,omitempty"`
	CEKWrapSenderIV       string `json:"cek_wrap_sender_iv,omitempty"`
	CEKWrapSenderSalt     string `json:"cek_wrap_sender_salt,omitempty"`
	ReplyTo               int    `json:"reply_to,omitempty"`
}
