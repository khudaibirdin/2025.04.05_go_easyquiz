package usecases

import (
	"app/internal/entities"
	"fmt"
)

type QuizUseCase struct {
	Repository QuizUseCaseRepository
}

type QuizUseCaseRepository interface {
	CreateQuiz(quiz entities.Quiz) (uint, error)
	CreateQuestions(quizID uint, questions []entities.Question) ([]uint, error)
	GetAllQuestions(quizID uint) ([]entities.Question, error)
	GetQuestion(quizID, number uint) (entities.Question, error)
	GetQuestionByNumber(quizID uint, lastNumber int) (entities.Question, error)
}

func NewQuizUseCase(r QuizUseCaseRepository) *QuizUseCase {
	return &QuizUseCase{Repository: r}
}

// Создание Квиза
func (uc *QuizUseCase) CreateQuiz(quiz entities.Quiz) (uint, error) {
	return uc.Repository.CreateQuiz(quiz)
}

// Создание вопросов для квиза
func (uc *QuizUseCase) CreateQuestions(quizID uint, questions []entities.Question) ([]uint, error) {
	return uc.Repository.CreateQuestions(quizID, questions)
}

// Получение всех вопросов по ID Квиза
func (uc *QuizUseCase) GetQuestions(quizID uint) ([]entities.Question, error) {
	return uc.Repository.GetAllQuestions(quizID)
}

// Получение вопроса из Квиза по ID или по номеру
func (uc *QuizUseCase) GetQuestion(quizID uint, questionID *uint, lastNumber *int) (entities.Question, error) {
	if lastNumber != nil {
		return uc.Repository.GetQuestionByNumber(quizID, *lastNumber+1)
	}
	if questionID != nil {
		return uc.Repository.GetQuestion(quizID, *questionID)
	}
	return entities.Question{}, fmt.Errorf("question get error")
}

// Получение количества вопросов в Квизе
func (uc *QuizUseCase) GetQuizQuestionsAmount(quizID uint) (int, error) {
	questions, err := uc.Repository.GetAllQuestions(quizID)
	if err != nil {
		return 0, err
	}
	return len(questions), nil
}
