package usecases

import (
	"app/internal/entities"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockQuizUseCaseRepository struct {
	mock.Mock
}

func (uc *MockQuizUseCaseRepository) CreateQuiz(quiz entities.Quiz) (uint, error) {
	return 0, nil
}
func (uc *MockQuizUseCaseRepository) CreateQuestions(quizID uint, questions []entities.Question) ([]uint, error) {
	return []uint{}, nil
}
func (uc *MockQuizUseCaseRepository) GetAllQuestions(quizID uint) ([]entities.Question, error) {
	return []entities.Question{}, nil
}
func (uc *MockQuizUseCaseRepository) GetQuestion(quizID, questionID uint) (entities.Question, error) {
	for _, question := range QUESTIONS {
		if question.ID == questionID {
			return question, nil
		}
	}
	return entities.Question{}, fmt.Errorf("no right question by ID")
}
func (uc *MockQuizUseCaseRepository) GetQuestionByNumber(quizID uint, number int) (entities.Question, error) {
	for _, question := range QUESTIONS {
		if question.Number == number {
			return question, nil
		}
	}
	return entities.Question{}, fmt.Errorf("no right question by number")
}

// Тест на получение вопроса по ID
func TestQuizGetQuestionByID(t *testing.T) {
	quizRepo := &MockQuizUseCaseRepository{}
	quizUseCase := NewQuizUseCase(quizRepo)
	var questionID uint = 1
	question, err := quizUseCase.GetQuestion(QUIZ_ID, &questionID, nil)
	if err != nil {
		t.Error("no quetion found")
	}
	if question.ID != questionID {
		t.Error("wrong quetionID returned")
	}

}

// Тест на получение следующего вопроса по его порядковому номеру в Квизе
func TestQuizGetQuestionByLastNumberSuccess(t *testing.T) {
	quizRepo := &MockQuizUseCaseRepository{}
	quizUseCase := NewQuizUseCase(quizRepo)
	var lastQuetionNumber int = 1
	question, err := quizUseCase.GetQuestion(QUIZ_ID, nil, &lastQuetionNumber)
	if err != nil {
		t.Error("no question found")
	}
	if question.Number != lastQuetionNumber+1 {
		t.Error("wrong quetion.Number returned")
	}
}

// Тест на получение следующего вопроса по его порядковому номеру в Квизе
//
// Если вопросы закончились, должна вернуться ошибка
func TestQuizGetQuestionByLastNumberEnd(t *testing.T) {
	quizRepo := &MockQuizUseCaseRepository{}
	quizUseCase := NewQuizUseCase(quizRepo)
	var lastQuestionNumber int = 2
	_, err := quizUseCase.GetQuestion(QUIZ_ID, nil, &lastQuestionNumber)
	if err == nil {
		t.Error("no question, but err == nil")
	}
}
