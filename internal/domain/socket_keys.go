package domain

type SocketKeys struct {
	ChatID         int    `json:"chat_id"`
	KyberPublicKey string `json:"kyber_public_key"`
	EcdhPublicKey  string `json:"ecdh_public_key"`
	EdPublicKey    string `json:"ed_public_key"`
}
