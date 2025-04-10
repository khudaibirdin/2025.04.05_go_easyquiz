package handlers

import (
	"fmt"
	"strconv"

	"app/internal/config"
	"app/internal/entities"
	"app/internal/usecases"

	"github.com/gofiber/fiber/v2"
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

// Создание квиза
func (h *QuizHandler) CreateQuiz(ctx *fiber.Ctx) error {
	type CreateQuizRequest struct {
		Theme string `json:"theme"`
	}
	var createQuizRequest CreateQuizRequest
	if err := ctx.BodyParser(&createQuizRequest); err != nil {
		return SetBadRequestResponse(
			ctx,
			false,
			fmt.Sprintf("Invalid request data, error: %s", err),
		)
	}
	userID := ctx.Locals("userID").(uint)
	quizID, err := h.UseCase.CreateQuiz(
		entities.Quiz{
			UserID: userID,
			Theme:  createQuizRequest.Theme,
		},
	)
	if err != nil {
		return SetBadRequestResponse(
			ctx,
			false,
			fmt.Sprintf("quiz creation error: %s", err),
		)
	}
	return SetSuccessResponse(
		ctx,
		true,
		quizID,
	)
}

// Начало теста
func (h *QuizHandler) StartQuiz(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(uint)
	quizIDstr := ctx.Params("quiz_id")
	quizID, err := strconv.Atoi(quizIDstr)
	if err != nil {
		SetBadRequestResponse(
			ctx,
			false,
			"quizID param parsing error",
		)
	}
	_, err = h.UseCase.StartQuiz(userID, uint(quizID))
	if err != nil {
		return SetBadRequestResponse(
			ctx,
			false,
			err,
		)
	}
	return SetSuccessResponse(
		ctx,
		true,
		nil,
	)
}

func (h *QuizHandler) CreateQuestion(ctx *fiber.Ctx) error {
	quizIDstr := ctx.Params("quiz_id")
	quizID, err := strconv.Atoi(quizIDstr)
	if err != nil {
		SetBadRequestResponse(
			ctx,
			false,
			"quizID param parsing error",
		)
	}
	type CreateQuestionRequest struct {
		Number int
		Text   string
	}
	var createQuestionRequest CreateQuestionRequest
	if err := ctx.BodyParser(&createQuestionRequest); err != nil {
		return SetBadRequestResponse(
			ctx,
			false,
			fmt.Sprintf("Invalid request data, error: %s", err),
		)
	}
	questionIDS, err := h.UseCase.CreateQuestions([]entities.Question{
		{
			QuizID: uint(quizID),
			Number: createQuestionRequest.Number,
			Text:   createQuestionRequest.Text,
		},
	})
	if err != nil {
		return SetBadRequestResponse(
			ctx,
			false,
			err,
		)
	}
	return SetSuccessResponse(
		ctx,
		true,
		questionIDS,
	)
}
