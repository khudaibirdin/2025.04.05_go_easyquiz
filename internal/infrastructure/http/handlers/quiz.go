package handlers

import (
	"app/internal/config"
	// "app/internal/entities"
	"app/internal/usecases"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type QuizHandler struct {
	UseCase usecases.QuizUseCase
	Config  *config.Config
}

func NewQuizHandler(uc usecases.QuizUseCase, cfg *config.Config) *QuizHandler {
	return &QuizHandler{
		UseCase: uc,
		Config:  cfg,
	}
}

func (h *QuizHandler) CreateQuiz(ctx *fiber.Ctx) error {
	type CreateQuizRequest struct {
		Theme string `json:"theme"`
	}
	var createQuizRequest CreateQuizRequest
	if err := ctx.BodyParser(&createQuizRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("Invalid request data, error: %s", err),
			},
		)
	}
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(uint)
	h.UseCase.CreateQuiz(
		entities.Quiz{
			User: 
		},
	)
	return ctx.SendStatus(fiber.StatusOK)
}
