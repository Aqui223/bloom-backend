package user

import (
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) IsUserWithEmailExists(c *fiber.Ctx) error {
	email := c.Query("email", "")
	if email == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"exists": false,
		})
	}

	_, err := h.userApp.IsUserWithEmailExists(email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"exists": false,
		})
	}

	return c.JSON(fiber.Map{
		"exists": true,
	})
}
