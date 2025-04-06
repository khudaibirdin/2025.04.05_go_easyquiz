package usecases

import (
	"app/internal/entities"
)

type QuizUseCase struct {
	Repository QuizUseCaseRepository
}

type QuizUseCaseRepository interface {
	CreateQuiz(quiz entities.Quiz) error
	CreateQuestions(quizID int, questions []entities.Question) error
	GetAllQuestions(quizID int) []entities.Question
	GetQuestion(quizID, number int) entities.Question
	GetQuestionByNumber(quizID, lastNumber int) entities.Question
}

func NewQuizUseCase(r QuizUseCaseRepository) *QuizUseCase {
	return &QuizUseCase{Repository: r}
}

func (uc *QuizUseCase) CreateQuiz(quiz entities.Quiz) {
	uc.Repository.CreateQuiz(quiz)
}

func (uc *QuizUseCase) CreateQuestions(quizID int, questions []entities.Question) error {
	uc.Repository.CreateQuestions(quizID, questions)
	return nil
}

func (uc *QuizUseCase) GetQuestions(quizID int) []entities.Question {
	return uc.Repository.GetAllQuestions(quizID)
}

func (uc *QuizUseCase) CheckQuestion(questionID int) entities.Question {
	return entities.Question{}
}

func (uc *QuizUseCase) GetQuestion(quizID, questionID, lastNumber *int) entities.Question {
	if lastNumber != nil {
		return uc.Repository.GetQuestionByNumber(*quizID, *lastNumber+1)
	}
	if questionID != nil {
		return uc.Repository.GetQuestion(*quizID, *questionID)
	}
	return entities.Question{}

}

func (uc *QuizUseCase) GetQuizQuestionsAmount(quizID int) int {
	return len(uc.Repository.GetAllQuestions(quizID))
}
