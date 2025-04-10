package usecases

import "app/internal/entities"

type ResultUseCase struct {
	Repository ResultRepository
}

type ResultRepository interface {
	Create(result entities.Result) (uint, error)
	Get(resultID uint) (*entities.Result, error)
}

func NewResultUseCase(r ResultRepository) *ResultUseCase {
	return &ResultUseCase{Repository: r}
}

// Создание результата теста для пользователя
//
// Вычисляется процент правильных ответов и заносится в БД
func (uc *ResultUseCase) Create(result entities.Result) (uint, error) {
	return uc.Repository.Create(result)
}
// result.Percent = int(float32(result.QuestionsAnsweredAmount) / float32(result.QuestionsAmount) * 100.0)
// Получить результат
func (uc *ResultUseCase) Get(result uint) (*entities.Result, error) {
	return uc.Repository.Get(result)
}
