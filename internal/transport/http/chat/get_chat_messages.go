package chat

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slipe-fun/skid-backend/internal/transport/http"
)

func (h *ChatHandler) GetChatMessages(c *fiber.Ctx) error {
	token, err := http.ExtractBearerToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	chatId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	messages, err := h.messageApp.GetChatMessages(token, chatId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if len(messages) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "no messages",
		})
	}

	return c.JSON(messages)
}
