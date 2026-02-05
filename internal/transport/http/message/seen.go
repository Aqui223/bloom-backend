package message

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/transport/http"
)

func (h *MessageHandler) Seen(c *fiber.Ctx) error {
	token, err := http.ExtractBearerToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "invalid_token",
			"message": "invalid token",
		})
	}

	var req struct {
		ChatID   int   `json:"chat_id"`
		Messages []int `json:"messages"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid_request",
			"message": "invalid request",
		})
	}

	validMessages, seenAt, chat, session, err := h.messageApp.UpdateMessagesSeenStatus(token, req.ChatID, req.Messages)
	if appErr, ok := err.(*domain.AppError); ok {
		return c.Status(appErr.Status).JSON(fiber.Map{
			"error":   appErr.Code,
			"message": appErr.Msg,
		})
	}

	outMsg := struct {
		Type     string    `json:"type"`
		UserID   int       `json:"user_id"`
		ChatID   int       `json:"chat_id"`
		SeenAt   time.Time `json:"seen_at"`
		Messages []int     `json:"messages"`
	}{
		Type:     "message.seen",
		UserID:   session.UserID,
		ChatID:   req.ChatID,
		SeenAt:   *seenAt,
		Messages: *validMessages,
	}

	b, err := json.Marshal(outMsg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "internal_error",
			"message": "internal error",
		})
	}

	h.wsHub.SendToUser(h.chatApp.GetOtherMember(chat, session.UserID).ID, b)

	return c.JSON(outMsg)
}
