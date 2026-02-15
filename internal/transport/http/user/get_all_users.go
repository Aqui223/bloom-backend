package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slipe-fun/skid-backend/internal/domain"
)

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 20)
	offset := c.QueryInt("offset", 0)

	if limit < 1 {
		limit = 20
	} else if limit > 20 {
		limit = 20
	}

	if offset < 0 {
		offset = 0
	}

	users, err := h.userApp.GetAllUsers(limit, offset)
	if appErr, ok := err.(*domain.AppError); ok {
		return c.Status(appErr.Status).JSON(fiber.Map{
			"error":   appErr.Code,
			"message": appErr.Msg,
		})
	}

	return c.JSON(users)
}
