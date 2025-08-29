package handler

import (
	"strconv"
	"strings"

	"github.com/gofiber/websocket/v2"
	"github.com/slipe-fun/skid-backend/internal/transport/ws/events"
	"github.com/slipe-fun/skid-backend/internal/transport/ws/types"
)

func HandleWS(hub *types.Hub) func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		defer c.Close()
		clientToken := c.Query("token")
		_, err := hub.JwtSvc.VerifyToken(clientToken)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte("Unauthorized"))
			c.Close()
			return
		}

		userID, err := hub.TokenSvc.ExtractUserID(clientToken)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte("Unauthorized"))
			c.Close()
			return
		}

		client := &types.Client{Conn: c}

		chats, err := hub.Chats.GetChatsByUserId(clientToken)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte("Get chats error"))
			c.Close()
			return
		}

		if len(chats) > 0 {
			for _, chat := range chats {
				events.Join(hub, client, "chat"+strconv.Itoa(chat.ID))
			}
		}

		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				events.Leave(hub, client)
				break
			}

			parts := strings.SplitN(string(msg), ":", 2)
			if len(parts) < 2 {
				continue
			}

			cmd := parts[0]
			data := parts[1]

			switch cmd {
			case "send":
				{
					roomMsg := strings.SplitN(data, "|", 2)
					if len(roomMsg) == 2 {
						events.Send(hub, client, userID, roomMsg[0], roomMsg[1])
					}
				}
			case "leave":
				events.Leave(hub, client)
			}
		}
	}
}
