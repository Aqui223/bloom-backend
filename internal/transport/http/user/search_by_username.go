package user

import "github.com/gofiber/fiber/v2"

func (h *UserHandler) SearchByUsername(c *fiber.Ctx) error {
	query := c.Query("q", "")
	limit := c.QueryInt("limit", 20)
	offset := c.QueryInt("offset", 0)

	if len(query) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "users_not_found",
		})
	}

	users, err := h.userApp.SearchUsersByUsername(query, limit, offset)
	if err != nil || len(users) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "users_not_found",
		})
	}

	return c.JSON(users)
}
