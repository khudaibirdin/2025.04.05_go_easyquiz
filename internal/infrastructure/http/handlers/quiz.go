package handlers

import (
	"app/internal/config"
	"app/internal/usecases"

	"github.com/gofiber/fiber/v2"
)

type QuizHandler struct {
	UseCase usecases.UserUsecase
	Config  *config.Config
}

func NewQuizHandler(uc usecases.UserUsecase, cfg *config.Config) *UserHandler {
	return &UserHandler{
		UseCase: uc,
		Config:  cfg,
	}
}

func (h *QuizHandler) CreateQuiz(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}
