package friend

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slipe-fun/skid-backend/internal/domain"
)

func (h *FriendHandler) DeleteFriend(c *fiber.Ctx) error {
	sessionVal := c.Locals("session")
	session, ok := sessionVal.(*domain.Session)
	if !ok {
		return fiber.ErrUnauthorized
	}

	var req struct {
		FriendID int `json:"friend_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid_request",
			"message": "invalid request",
		})
	}

	if req.FriendID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "no_friend_id",
			"message": "no friend_id",
		})
	}

	err := h.friendApp.DeleteFriend(session.UserID, req.FriendID)
	if appErr, ok := err.(*domain.AppError); ok {
		return c.Status(appErr.Status).JSON(fiber.Map{
			"error":   appErr.Code,
			"message": appErr.Msg,
		})
	}

	return c.JSON(fiber.Map{
		"message": "friend deleted",
	})
}
