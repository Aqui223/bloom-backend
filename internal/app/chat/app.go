package chat

type ChatApp struct {
	chats ChatRepo
}

func NewChatApp(chats ChatRepo) *ChatApp {
	return &ChatApp{
		chats: chats,
	}
}
