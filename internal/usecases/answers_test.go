package usecases

import (
	"app/internal/entities"
	"testing"
)

var (
	USER_ID             int = 1
	QUIZ_ID             int = 1
	QUIZ_ANSWERS_AMOUNT int = 2
	QUESTIONS               = []entities.Question{
		{
			ID:      1,
			QuizID:  QUIZ_ID,
			Number:  1,
			Text:    "Сколько пальцев на руке здорового человека?",
			Answers: []string{"5", "1", "2", "7"},
			Right:   1,
		},
		{
			ID:      2,
			QuizID:  QUIZ_ID,
			Number:  2,
			Text:    "Из чего делают окна?",
			Answers: []string{"Пластик", "Стекло", "Полиуретан", "Бумага"},
			Right:   2,
		},
	}
)

type MockAnswerRepositiry struct{}

func (uc *MockAnswerRepositiry) Create(userID, quizID, questionID, answer int) error {
	return nil
}

func (uc *MockAnswerRepositiry) GetAll(userID, quizID int) []entities.Answers {
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
	return answers
}

type MockQuizRepositiry struct{}

func (uc *MockQuizRepositiry) CreateQuiz(quiz entities.Quiz) error {
	return nil
}

func (uc *MockQuizRepositiry) CreateQuestions(quizID int, questions []entities.Question) error {
	return nil
}

func (uc *MockQuizRepositiry) GetAllQuestions(quizID int) []entities.Question {
	return QUESTIONS
}

func (uc *MockQuizRepositiry) GetQuestion(quizID, questionID int) entities.Question {
	for _, question := range QUESTIONS {
		if question.ID == questionID {
			return question
		}
	}
	return entities.Question{}
}

func (uc *MockQuizRepositiry) GetQuestionByNumber(quizID, lastNumber int) entities.Question {
	return entities.Question{}
}

type MockResultRepository struct{}

func (uc *MockResultRepository) Create(result entities.Result) entities.Result {
	return result
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
	result := answerUseCase.CheckAll(USER_ID, QUIZ_ID)
	if result.Percent != 50 {
		t.Errorf("test result != 50")
	}
}
