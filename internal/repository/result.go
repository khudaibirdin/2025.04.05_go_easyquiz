package repository

import (
	"app/internal/entities"
	"errors"

	"gorm.io/gorm"
)

type ResultRepository struct {
	db *gorm.DB
}

func NewResultRepository(db *gorm.DB) *ResultRepository {
	return &ResultRepository{
		db: db,
	}
}

func (r *ResultRepository) Create(newResult entities.Result) (uint, error) {
	result := r.db.Create(&newResult)
	return newResult.ID, result.Error
}
func (r *ResultRepository) Get(resultID uint) (*entities.Result, error) {
	result := &entities.Result{}
	res := r.db.First(result, resultID)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return result, res.Error
}
