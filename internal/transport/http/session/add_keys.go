package session

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slipe-fun/skid-backend/internal/domain"
)

func (h *SessionHandler) AddKeys(c *fiber.Ctx) error {
	sessionVal := c.Locals("session")
	session, ok := sessionVal.(*domain.Session)
	if !ok {
		return fiber.ErrUnauthorized
	}

	var req struct {
		IdentityPublicKey string `json:"identity_pub"`
		EcdhPublicKey     string `json:"ecdh_pub"`
		KyberPublicKey    string `json:"kyber_pub"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid_request",
			"message": "invalid request",
		})
	}

	if req.IdentityPublicKey == "" || req.EcdhPublicKey == "" || req.KyberPublicKey == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "not_all_keys",
			"message": "not all keys are provided",
		})
	}

	err := h.sessionApp.AddKeys(session.ID, session.UserID, req.IdentityPublicKey, req.EcdhPublicKey, req.KyberPublicKey)
	if appErr, ok := err.(*domain.AppError); ok {
		return c.Status(appErr.Status).JSON(fiber.Map{
			"error":   appErr.Code,
			"message": appErr.Msg,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
