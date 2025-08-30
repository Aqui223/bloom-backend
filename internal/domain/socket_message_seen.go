package domain

type SocketMessageSeen struct {
	ChatID   int   `json:"chat_id"`
	Messages []int `json:"messages"`
}
