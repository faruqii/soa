package handler

import (
	"github.com/faruqii/soa/authservice/internal/dto"
	"github.com/faruqii/soa/authservice/internal/service"
	"github.com/gofiber/fiber/v2"
)

type KeyHandler struct {
	svc service.KeyService
}

func NewKeyHandler(svc service.KeyService) *KeyHandler {
	return &KeyHandler{
		svc: svc,
	}
}

func (h *KeyHandler) CreateKey(ctx *fiber.Ctx) error {
	var req dto.APIKey
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if req.UserID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}

	apikey, err := h.svc.CreateKey(req.UserID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "API key created successfully",
		"key":     apikey.Key,
		"user_id": apikey.UserID,
	})
}

func (h *KeyHandler) ValidateKey(ctx *fiber.Ctx) error {
	key := ctx.Params("key")
	if key == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Key is required",
		})
	}

	apikey, err := h.svc.ValidateKey(key)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid API key",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "API key is valid",
		"key":     apikey.Key,
	})
}
