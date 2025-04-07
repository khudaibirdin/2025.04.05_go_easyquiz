package usecases

import (
	"app/internal/entities"
	"fmt"
	"testing"

	"gorm.io/gorm"
)

var (
	USER_ID             uint = 1
	QUIZ_ID             uint = 1
	QUIZ_ANSWERS_AMOUNT int  = 2
	QUESTIONS                = []entities.Question{
		{
			Model:   gorm.Model{ID: 1},
			QuizID:  QUIZ_ID,
			Number:  1,
			Text:    "Сколько пальцев на руке здорового человека?",
			Answers: []string{"5", "1", "2", "7"},
			Right:   1,
		},
		{
			Model:   gorm.Model{ID: 2},
			QuizID:  QUIZ_ID,
			Number:  2,
			Text:    "Из чего делают окна?",
			Answers: []string{"Пластик", "Стекло", "Полиуретан", "Бумага"},
			Right:   2,
		},
	}
)

type MockAnswerRepositiry struct{}

func (uc *MockAnswerRepositiry) Create(userID, quizID, questionID uint, answer int) (uint, error) {
	return 0, nil
}

func (uc *MockAnswerRepositiry) GetAll(userID, quizID uint) ([]entities.Answers, error) {
	answers := []entities.Answers{
		{
			UserID:     USER_ID,
			QuizID:     QUIZ_ID,
			QuestionID: 1,
			Answer:     1,
		},
		{
			UserID:     USER_ID,
			QuizID:     QUIZ_ID,
			QuestionID: 2,
			Answer:     3,
		},
	}
	return answers, nil
}

type MockQuizRepositiry struct{}

func (uc *MockQuizRepositiry) CreateQuiz(quiz entities.Quiz) (uint, error) {
	return 0, nil
}

func (uc *MockQuizRepositiry) CreateQuestions(quizID uint, questions []entities.Question) ([]uint, error) {
	return []uint{}, nil
}

func (uc *MockQuizRepositiry) GetAllQuestions(quizID uint) ([]entities.Question, error) {
	return QUESTIONS, nil
}

func (uc *MockQuizRepositiry) GetQuestion(quizID, questionID uint) (entities.Question, error) {
	for _, question := range QUESTIONS {
		if question.ID == questionID {
			return question, nil
		}
	}
	return entities.Question{}, fmt.Errorf("no question found")
}

func (uc *MockQuizRepositiry) GetQuestionByNumber(quizID uint, lastNumber int) (entities.Question, error) {
	return entities.Question{}, nil
}

type MockResultRepository struct{
	result entities.Result
}

func (uc *MockResultRepository) Create(result entities.Result) (uint, error) {
	uc.result = result
	return uc.result.ID, nil
}

func (uc *MockResultRepository) Get(resultID uint) (entities.Result, error) {
	return uc.result, nil
}

// Тест кейс на проверку количесттва правильных ответов в квизе для пользователя
//
// Результат выводит также процент правильности
func TestCheckAll(t *testing.T) {
	quizRepo := MockQuizRepositiry{}
	quizUseCase := NewQuizUseCase(&quizRepo)
	resultRepo := MockResultRepository{}
	resultUseCase := NewResultUseCase(&resultRepo)
	answerRepo := MockAnswerRepositiry{}
	answerUseCase := NewAnswersUseCase(&answerRepo, quizUseCase, resultUseCase)
	result, err := answerUseCase.CheckAll(USER_ID, QUIZ_ID)
	if err != nil {
		t.Error("no result returned")
	}
	if result.Percent != 50 {
		t.Error("test result != 50")
	}
}
