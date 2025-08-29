package chat

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/slipe-fun/skid-backend/internal/transport/http"
)

func (h *ChatHandler) AddChatKeys(c *fiber.Ctx) error {
	token, err := http.ExtractBearerToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	chatId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if chatId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "chat id required"})
	}

	var req struct {
		KyberPublicKey string `json:"kyberPublicKey"`
		EcdhPublicKey  string `json:"ecdhPublicKey"`
		EdPublicKey    string `json:"edPublicKey"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if req.KyberPublicKey == "" || req.EcdhPublicKey == "" || req.EdPublicKey == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "u need to specify all keys"})
	}

	kyberKey, err := base64.StdEncoding.DecodeString(req.KyberPublicKey)
	if err != nil {
		return errors.New("invalid base64 for Kyber key")
	}
	ecdhKey, err := base64.StdEncoding.DecodeString(req.EcdhPublicKey)
	if err != nil {
		return errors.New("invalid base64 for ECDH key")
	}
	edKey, err := base64.StdEncoding.DecodeString(req.EdPublicKey)
	if err != nil {
		return errors.New("invalid base64 for Ed25519 key")
	}

	if len(kyberKey) != 1184 {
		return errors.New("invalid Kyber key length")
	}
	fmt.Print(len(ecdhKey))
	if len(ecdhKey) != 44 {
		return errors.New("invalid ECDH key length")
	}
	if len(edKey) != 44 {
		return errors.New("invalid Ed25519 key length")
	}

	chat, err := h.chatApp.GetChatById(token, chatId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updateChatErr := h.chatApp.AddKeys(token, chat, req.KyberPublicKey, req.EcdhPublicKey, req.EdPublicKey)
	if updateChatErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": updateChatErr.Error()})
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
