package chat

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slipe-fun/skid-backend/internal/transport/http"
)

func (h *ChatHandler) GetChatById(c *fiber.Ctx) error {
	token, err := http.ExtractBearerToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid_token",
		})
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid_params",
		})
	}

	chat, err := h.chatApp.GetChatById(token, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "chat_not_found",
		})
	}

	return c.JSON(fiber.Map{
		"id":             chat.ID,
		"members":        chat.Members,
		"encryption_key": chat.EncryptionKey,
	})
}
