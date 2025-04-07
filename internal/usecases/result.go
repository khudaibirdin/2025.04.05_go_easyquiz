package usecases

import "app/internal/entities"

type ResultUseCase struct {
	Repository ResultUseCaseRepository
}

type ResultUseCaseRepository interface {
	Create(result entities.Result) (uint, error)
	Get(resultID uint) (entities.Result, error)
}

func NewResultUseCase(r ResultUseCaseRepository) *ResultUseCase {
	return &ResultUseCase{Repository: r}
}

// Создание результата теста для пользователя
//
// Вычисляется процент правильных ответов и заносится в БД
func (uc *ResultUseCase) Create(result entities.Result) (uint, error) {
	result.Percent = int(float32(result.QuestionsAnsweredAmount) / float32(result.QuestionsAmount) * 100.0)
	return uc.Repository.Create(result)
}

// Получить результат
func (uc *ResultUseCase) Get(result uint) (entities.Result, error) {
	return uc.Repository.Get(result)
}