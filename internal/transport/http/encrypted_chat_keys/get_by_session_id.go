package encryptedchatkeys

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slipe-fun/skid-backend/internal/domain"
)

func (h *EncryptedChatKeysHandler) GetBySessionID(c *fiber.Ctx) error {
	sessionVal := c.Locals("session")
	session, ok := sessionVal.(*domain.Session)
	if !ok {
		return fiber.ErrUnauthorized
	}

	keys, err := h.keys.GetBySessionID(session.ID)
	if appErr, ok := err.(*domain.AppError); ok {
		return c.Status(appErr.Status).JSON(fiber.Map{
			"error":   appErr.Code,
			"message": appErr.Msg,
		})
	}

	return c.JSON(keys)
}
