package domain

import "time"

type Message struct {
	ID                    int        `db:"id" json:"id"`
	Ciphertext            string     `db:"ciphertext" json:"ciphertext"`
	Nonce                 string     `db:"nonce" json:"nonce"`
	ChatID                int        `db:"chat_id" json:"chat_id"`
	EncapsulatedKey       *string    `db:"encapsulated_key" json:"encapsulated_key,omitempty"`
	Signature             *string    `db:"signature" json:"signature,omitempty"`
	SignedPayload         *string    `db:"signed_payload" json:"signed_payload,omitempty"`
	CEKWrap               *string    `db:"cek_wrap" json:"cek_wrap,omitempty"`
	CEKWrapIV             *string    `db:"cek_wrap_iv" json:"cek_wrap_iv,omitempty"`
	CEKWrapSalt           *string    `db:"cek_wrap_salt" json:"cek_wrap_salt,omitempty"`
	EncapsulatedKeySender *string    `db:"encapsulated_key_sender" json:"encapsulated_key_sender,omitempty"`
	CEKWrapSender         *string    `db:"cek_wrap_sender" json:"cek_wrap_sender,omitempty"`
	CEKWrapSenderIV       *string    `db:"cek_wrap_sender_iv" json:"cek_wrap_sender_iv,omitempty"`
	CEKWrapSenderSalt     *string    `db:"cek_wrap_sender_salt" json:"cek_wrap_sender_salt,omitempty"`
	Seen                  *time.Time `db:"seen" json:"seen,omitempty"`
	ReplyTo               *int       `db:"reply_to" json:"reply_to,omitempty"`
}
