package usecases

import "app/internal/entities"

type ResultUseCase struct {
	Repository ResultUseCaseRepository
}

type ResultUseCaseRepository interface {
	Create(result entities.Result) entities.Result
}

func NewResultUseCase(r ResultUseCaseRepository) *ResultUseCase {
	return &ResultUseCase{Repository: r}
}

func (uc *ResultUseCase) CreateResult(result entities.Result) entities.Result {
	result.Percent = int(float32(result.QuestionsAnsweredAmount) / float32(result.QuestionsAmount) * 100.0)
	return uc.Repository.Create(result)
}
