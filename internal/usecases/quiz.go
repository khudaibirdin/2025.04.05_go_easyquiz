package usecases

import (
	"app/internal/entities"
	"fmt"
)

type QuizUseCase struct {
	Repository     QuizUseCaseRepository
	ResultUseCase  ResultUseCase
	// AnswersUseCase AnswersUseCase
}

type QuizUseCaseRepository interface {
	// Квиз
	CreateQuiz(quiz entities.Quiz) (uint, error)

	// Вопросы для квиза
	CreateQuestions(questions []entities.Question) ([]uint, error)
	GetQuestion(quizID, questionID uint) (*entities.Question, error)
	GetQuestionByNumber(quizID uint, number int) (*entities.Question, error)
	GetAllQuestions(quizID uint) (*[]entities.Question, error)

	// Варианты ответов для вопроса
	CreateAnswerVariant(answerVariant entities.AnswerVariant) (uint, error)
	GetAnswerVariant(answerVariantID uint) (*entities.AnswerVariant, error)
	GetQuestionAnswerVariants(quizID uint) (*[]entities.AnswerVariant, error)
}

func NewQuizUseCase(r QuizUseCaseRepository, resultUseCase ResultUseCase) *QuizUseCase {
	return &QuizUseCase{
		Repository:     r,
		ResultUseCase:  resultUseCase,
		// AnswersUseCase: answersUseCase,
	}
}

// Создание Квиза
func (uc *QuizUseCase) CreateQuiz(quiz entities.Quiz) (uint, error) {
	return uc.Repository.CreateQuiz(quiz)
}

// Создание вопросов для квиза
func (uc *QuizUseCase) CreateQuestions(questions []entities.Question) ([]uint, error) {
	return uc.Repository.CreateQuestions(questions)
}

// Получение всех вопросов по ID Квиза
func (uc *QuizUseCase) GetQuestions(quizID uint) (*[]entities.Question, error) {
	return uc.Repository.GetAllQuestions(quizID)
}

// Получение вопроса из Квиза по ID или по номеру
func (uc *QuizUseCase) GetQuestion(quizID uint, questionID *uint, lastNumber *int) (*entities.Question, error) {
	if lastNumber != nil {
		return uc.Repository.GetQuestionByNumber(quizID, *lastNumber+1)
	}
	if questionID != nil {
		return uc.Repository.GetQuestion(quizID, *questionID)
	}
	return nil, fmt.Errorf("question get error")
}

// Получение количества вопросов в Квизе
func (uc *QuizUseCase) GetQuizQuestionsAmount(quizID uint) (int, error) {
	questions, err := uc.Repository.GetAllQuestions(quizID)
	if err != nil {
		return 0, err
	}
	return len(*questions), nil
}

func (uc *QuizUseCase) StartQuiz(userID, quizID uint) (uint, error) {
	questions, err := uc.GetQuestions(quizID)
	if err != nil {
		return 0, err
	}
	return uc.ResultUseCase.Create(
		entities.Result{
			UserID:          userID,
			QuizID:          quizID,
			QuestionsAmount: len(*questions),
		},
	)
}

// Создание варианта ответа для вопроса
func (uc *QuizUseCase) CreateAnswerVariant(answerVariant entities.AnswerVariant) (uint, error) {
	return uc.Repository.CreateAnswerVariant(answerVariant)
}

// получение варианта ответа для вопроса по id
func (uc *QuizUseCase) GetAnswerVariant(answerVariant uint) (*entities.AnswerVariant, error) {
	return uc.Repository.GetAnswerVariant(answerVariant)
}

// получение всех вариантов ответа для вопроса по id квиза
func (uc *QuizUseCase) GetQuestionAnswerVariants(questionID uint) (*[]entities.AnswerVariant, error) {
	return uc.Repository.GetQuestionAnswerVariants(questionID)
}
