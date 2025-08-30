package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slipe-fun/skid-backend/internal/config"
)

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid_request"})
	}

	if req.Password == "" || req.Username == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid_request"})
	}

	if len(req.Username) < 4 || !config.UsernameRegex.MatchString(req.Username) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid_username"})
	}

	token, user, err := h.authApp.Register(req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "cant_create_user"})
	}

	return c.JSON(fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"date":     user.Date,
		},
	})
}
