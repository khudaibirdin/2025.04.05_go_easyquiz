package usecases

import "app/internal/entities"

type AnswersUseCase struct {
	Repository    AnswersUseCaseRepository
	QuizUseCase   *QuizUseCase
	ResultUsecase *ResultUseCase
}

type AnswersUseCaseRepository interface {
	Create(userID, quizID, questionID uint, answer int) (uint, error)
	GetAll(userID, quizID uint) ([]entities.Answers, error)
}

func NewAnswersUseCase(r AnswersUseCaseRepository, quizUseCase *QuizUseCase, resultUseCase *ResultUseCase) *AnswersUseCase {
	return &AnswersUseCase{
		Repository:    r,
		QuizUseCase:   quizUseCase,
		ResultUsecase: resultUseCase,
	}
}

// Регистрация ответа пользователя на вопрос Квиза
func (uc *AnswersUseCase) Register(userID, quizID, questionID uint, answer int) (uint, error) {
	return uc.Repository.Create(userID, quizID, questionID, answer)
}

// Проверка ответов на Квиз для конкретного пользователя с сохранением результата
func (uc *AnswersUseCase) CheckAll(userID, quizID uint) (entities.Result, error) {
	userAnswers, err := uc.Repository.GetAll(userID, quizID)
	if err != nil {
		return entities.Result{}, err
	}
	var (
		questionsAmount         int
		questionsAnsweredAmount int
	)
	questionsAmount, err = uc.QuizUseCase.GetQuizQuestionsAmount(quizID)
	if err != nil {
		return entities.Result{}, nil
	}
	for _, userAnswer := range userAnswers {
		rightAnswer, err := uc.QuizUseCase.GetQuestion(userAnswer.QuizID, &userAnswer.QuestionID, nil)
		if err != nil {
			return entities.Result{}, err
		}
		if userAnswer.Answer == rightAnswer.Right {
			questionsAnsweredAmount++
		}
	}
	resultID, err := uc.ResultUsecase.Create(
		entities.Result{
			UserID:                  userID,
			QuizID:                  quizID,
			QuestionsAmount:         questionsAmount,
			QuestionsAnsweredAmount: questionsAnsweredAmount,
		},
	)
	if err != nil {
		return entities.Result{}, err
	}
	return uc.ResultUsecase.Repository.Get(resultID)
}
