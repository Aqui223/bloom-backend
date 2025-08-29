package user

import "github.com/gofiber/fiber/v2"

func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := h.userApp.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"id":       user.ID,
		"username": user.Username,
		"date":     user.Date,
	})
}
