package repository

import (
	"app/internal/entities"

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

func (r *UserRepository) Get(string) (entities.User, error) {
	return entities.User{}, nil
}
