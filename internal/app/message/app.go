package message

type MessageApp struct {
	messages MessageRepo
	chats    ChatApp
}

func NewMessageApp(messages MessageRepo,
	chats ChatApp,
) *MessageApp {
	return &MessageApp{
		messages: messages,
		chats:    chats,
	}
}
