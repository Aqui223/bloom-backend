package chat

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slipe-fun/skid-backend/internal/transport/http"
)

func (h *ChatHandler) CreateChat(c *fiber.Ctx) error {
	token, err := http.ExtractBearerToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var req struct {
		Recipient int `json:"recipient"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if req.Recipient == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "recipient required"})
	}

	user, err := h.userApp.GetUserById(req.Recipient)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "recipient not found",
		})
	}

	chat1, err := h.chatApp.GetChatWithUsers(token, req.Recipient)
	if chat1 != nil || err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "chat with user already exists",
		})
	}

	chat, err := h.chatApp.CreateChat(token, user.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"id":      chat.ID,
		"members": chat.Members,
	})
}
