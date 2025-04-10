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

func (r *UserRepository) GetByUserName(username string) (*entities.User, error) {
	user := &entities.User{}
	result := r.db.Where(&entities.User{Login: username}).First(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, result.Error
}

func (r *UserRepository) Get(userID uint) (*entities.User, error) {
	user := &entities.User{}
	result := r.db.First(user, userID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, result.Error
}
