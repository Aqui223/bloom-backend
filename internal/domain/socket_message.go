package domain

type SocketMessage struct {
	Ciphertext            string `json:"ciphertext"`
	EncapsulatedKey       string `json:"encapsulated_key"`
	Nonce                 string `json:"nonce"`
	ChatID                int    `json:"chat_id"`
	Signature             string `json:"signature"`
	SignedPayload         string `json:"signed_payload"`
	CEKWrap               string `json:"cek_wrap"`
	CEKWrapIV             string `json:"cek_wrap_iv"`
	CEKWrapSalt           string `json:"cek_wrap_salt"`
	EncapsulatedKeySender string `json:"encapsulated_key_sender"`
	CEKWrapSender         string `json:"cek_wrap_sender"`
	CEKWrapSenderIV       string `json:"cek_wrap_sender_iv"`
	CEKWrapSenderSalt     string `json:"cek_wrap_sender_salt"`
}
