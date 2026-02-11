package friend

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slipe-fun/skid-backend/internal/domain"
)

func (h *FriendHandler) GetFriends(c *fiber.Ctx) error {
	sessionVal := c.Locals("session")
	session, ok := sessionVal.(*domain.Session)
	if !ok {
		return fiber.ErrUnauthorized
	}

	limit := c.QueryInt("limit", 20)
	offset := c.QueryInt("offset", 0)
	status := c.Params("status", "accepted")

	if status != "pending" && status != "accepted" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid_status",
			"message": "invalid status",
		})
	}

	friends, err := h.friendApp.GetFriends(session.UserID, status, limit, offset)
	if appErr, ok := err.(*domain.AppError); ok {
		return c.Status(appErr.Status).JSON(fiber.Map{
			"error":   appErr.Code,
			"message": appErr.Msg,
		})
	}

	return c.JSON(friends)
}
