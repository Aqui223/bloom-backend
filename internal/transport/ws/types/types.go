package types

import (
	"github.com/gofiber/websocket/v2"
	ChatApp "github.com/slipe-fun/skid-backend/internal/app/chat"
	MessageApp "github.com/slipe-fun/skid-backend/internal/app/message"
	UserApp "github.com/slipe-fun/skid-backend/internal/app/user"
	"github.com/slipe-fun/skid-backend/internal/service"
)

type Client struct {
	Conn  *websocket.Conn
	Rooms map[string]bool
}

type Hub struct {
	Clients         map[string]map[*Client]bool
	ClientsByUserID map[int]*Client
	Chats           *ChatApp.ChatApp
	Messages        *MessageApp.MessageApp
	Users           *UserApp.UserApp
	JwtSvc          *service.JWTService
	TokenSvc        *service.TokenService
}

func NewHub(Chats *ChatApp.ChatApp, Messages *MessageApp.MessageApp, Users *UserApp.UserApp, JwtSvc *service.JWTService, TokenSvc *service.TokenService) *Hub {
	return &Hub{
		Clients:  make(map[string]map[*Client]bool),
		Chats:    Chats,
		Messages: Messages,
		Users:    Users,
		JwtSvc:   JwtSvc,
		TokenSvc: TokenSvc,
	}
}

func (h *Hub) JoinRoom(client *Client, room string) {
	if client.Rooms == nil {
		client.Rooms = make(map[string]bool)
	}
	client.Rooms[room] = true

	if _, ok := h.Clients[room]; !ok {
		h.Clients[room] = make(map[*Client]bool)
	}
	h.Clients[room][client] = true
}

func (h *Hub) LeaveRoom(client *Client, room string) {
	if _, exists := client.Rooms[room]; exists {
		delete(client.Rooms, room)
	}

	if clients, ok := h.Clients[room]; ok {
		delete(clients, client)
		if len(clients) == 0 {
			delete(h.Clients, room)
		}
	}
}

func (h *Hub) LeaveAllRooms(client *Client) {
	for room := range client.Rooms {
		h.LeaveRoom(client, room)
	}
}

func (h *Hub) Broadcast(room string, message []byte) {
	if clients, ok := h.Clients[room]; ok {
		for client := range clients {
			client.Conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}
