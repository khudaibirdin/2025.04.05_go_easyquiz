package usecases

import "app/internal/entities"

type AnswersUseCase struct {
	Repository    AnswersUseCaseRepository
	QuizUseCase   *QuizUseCase
	ResultUsecase *ResultUseCase
}

type AnswersUseCaseRepository interface {
	Create(userID, quizID, questionID, answer int) error
	GetAll(userID, quizID int) []entities.Answers
}

func NewAnswersUseCase(r AnswersUseCaseRepository, quizUseCase *QuizUseCase, resultUseCase *ResultUseCase) *AnswersUseCase {
	return &AnswersUseCase{
		Repository:  r,
		QuizUseCase: quizUseCase,
		ResultUsecase: resultUseCase,
	}
}

func (uc *AnswersUseCase) Register(userID, quizID, questionID, answer int) error {
	return uc.Repository.Create(userID, quizID, questionID, answer)
}

func (uc *AnswersUseCase) CheckAll(userID, quizID int) entities.Result {
	answers := uc.Repository.GetAll(userID, quizID)
	var (
		questionsAmount         int
		questionsAnsweredAmount int
	)
	questionsAmount = uc.QuizUseCase.GetQuizQuestionsAmount(quizID)
	for _, answer := range answers {
		if answer.Answer == uc.QuizUseCase.GetQuestion(&answer.QuizID, &answer.QuestionID, nil).Right {
			questionsAnsweredAmount++
		}
	}
	return uc.ResultUsecase.CreateResult(
		entities.Result{
			UserID:                  userID,
			QuizID:                  quizID,
			QuestionsAmount:         questionsAmount,
			QuestionsAnsweredAmount: questionsAnsweredAmount,
		},
	)
}
