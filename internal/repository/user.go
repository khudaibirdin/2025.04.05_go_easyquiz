package repository

import (
	"app/internal/entities"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Add(newUser entities.User) (uint, error) {
	result := r.db.Create(&newUser)
	return newUser.ID, result.Error
}

func (r *UserRepository) Get(string) (*entities.User, error) {
	user := &entities.User{}
	result := r.db.First(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, result.Error
}
