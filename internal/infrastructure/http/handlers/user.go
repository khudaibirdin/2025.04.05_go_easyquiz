package handlers

import (
	"app/internal/config"
	"app/internal/usecases"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UseCase usecases.UserUsecase
	Config  *config.Config
}

func NewUserHandler(uc usecases.UserUsecase, cfg *config.Config) *UserHandler {
	return &UserHandler{
		UseCase: uc,
		Config:  cfg,
	}
}

// Хэндлер Логина
func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	type LoginRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	var req LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("Invalid request data, error: %s", err),
			},
		)
	}

	user, err := h.UseCase.Login(req.Login, req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("Error while login, error: %s", err),
			},
		)
	}
	token, err := Token{User: &user}.Generate()
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	
	return ctx.JSON(
		fiber.Map{
			"success": true,
			"error":   nil,
			"jwt":     token,
		},
	)
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req usecases.UserRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}
	userID, err := h.UseCase.Register(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"success": false,
				"error":   err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"success": true,
			"id":      userID,
		},
	)
}
