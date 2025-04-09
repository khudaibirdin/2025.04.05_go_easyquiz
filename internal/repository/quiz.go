package repository

type NewQuizRepository struct {
	Repository NewQuizRepositoryRepository
}

type NewQuizRepositoryRepository interface {
}

func NewNewQuizRepository(r NewQuizRepositoryRepository) *NewQuizRepository {
	return &NewQuizRepository{Repository: r}
}

