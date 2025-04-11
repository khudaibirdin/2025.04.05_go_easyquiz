package handlers

import (
	"fmt"

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
	quizID, err := ctx.ParamsInt("quiz_id")
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

// Создание вопроса
func (h *QuizHandler) CreateQuestion(ctx *fiber.Ctx) error {
	quizID, err := ctx.ParamsInt("quiz_id")
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

type CreateAnswerVariantRequest struct {
	Text    string `json:"text"`
	IsRight bool   `json:"is_right"`
}

// создание варианта ответа для вопроса
func (h *QuizHandler) CreateAnswerVariant(ctx *fiber.Ctx) error {
	questionID, err := ctx.ParamsInt("question_id")
	if err != nil {
		return SetBadRequestResponse(
			ctx,
			false,
			err,
		)
	}
	var createAnswerVariantRequest CreateAnswerVariantRequest
	if err := ctx.BodyParser(&createAnswerVariantRequest); err != nil {
		return SetBadRequestResponse(
			ctx,
			false,
			fmt.Sprintf("Invalid request data, error: %s", err),
		)
	}
	answerVariantID, err := h.UseCase.CreateAnswerVariant(
		entities.AnswerVariant{
			QuestionID: uint(questionID),
			Text:       createAnswerVariantRequest.Text,
			IsRight:    createAnswerVariantRequest.IsRight,
		},
	)
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
		answerVariantID,
	)
}

type GetQuestionAnswersResponse struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

// Получение всех вариантов ответа для вопроса
func (h *QuizHandler) GetQuestionAnswers(ctx *fiber.Ctx) error {
	questionID, err := ctx.ParamsInt("question_id")
	if err != nil {
		return SetBadRequestResponse(
			ctx,
			false,
			err,
		)
	}
	answerVariants, err := h.UseCase.GetQuestionAnswerVariants(uint(questionID))
	if err != nil {
		return SetBadRequestResponse(
			ctx,
			false,
			err,
		)
	}
	var getQuestionAnswersResponse []GetQuestionAnswersResponse
	for _, answerVariant := range *answerVariants {
		getQuestionAnswersResponse = append(getQuestionAnswersResponse, GetQuestionAnswersResponse{answerVariant.ID, answerVariant.Text})
	}
	return SetSuccessResponse(
		ctx,
		true,
		getQuestionAnswersResponse,
	)
}
