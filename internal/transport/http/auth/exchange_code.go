package auth

import (
	"github.com/gofiber/fiber/v2"
)

func (h *AuthHandler) ExchangeCode(c *fiber.Ctx) error {
	if c.Query("state") != "random-state" {
		return c.Status(400).SendString("invalid state")
	}

	code := c.Query("code")
	if code == "" {
		return c.Status(400).SendString("no code")
	}

	token, user, err := h.authApp.ExchangeCode(code)
	if err != nil {
		return c.Status(400).SendString("server error")
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
