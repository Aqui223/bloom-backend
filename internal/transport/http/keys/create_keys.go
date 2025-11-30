package keys

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/transport/http"
)

func (h *KeysHandler) SaveChatKeys(c *fiber.Ctx) error {
	token, err := http.ExtractBearerToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid_token",
		})
	}

	chatId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid_params"})
	}

	if chatId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "no_chat"})
	}

	var req struct {
		Ciphertext string `json:"ciphertext"`
		Nonce      string `json:"nonce"`
		Salt       string `json:"salt"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid_request"})
	}

	if req.Ciphertext == "" || req.Nonce == "" || req.Salt == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid_request"})
	}

	chat, err := h.chatApp.GetChatById(token, chatId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "chat_not_found"})
	}

	_, err = h.keysApp.CreateKeys(token, chat.ID, &domain.EncryptedKeys{
		Ciphertext: req.Ciphertext,
		Nonce:      req.Nonce,
		Salt:       req.Salt,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cant_save_chat_keys"})
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
