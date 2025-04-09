package usecases

import (
	"app/internal/entities"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Add(entities.User) (uint, error)
	Get(username string) (*entities.User, error)
}

type UserUsecase struct {
	Repository UserRepository
}

func NewUserUsecase(r UserRepository) *UserUsecase {
	return &UserUsecase{Repository: r}
}

type UserRegisterRequest struct {
	Login    string
	Password string
}

// Регистрация пользователя
func (uc *UserUsecase) Register(user UserRegisterRequest) error {
	if len(user.Password) < 8 {
		return fmt.Errorf("password's len is less than 8")
	}
	userExists, err := uc.Repository.Get(user.Login)
	if err != nil {
		return fmt.Errorf("database error: %s", err)
	}
	if userExists != nil {
		return fmt.Errorf("user is already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return fmt.Errorf("failed to hash password: %w", err)
    }
	_, err = uc.Repository.Add(
		entities.User{
			Login:    user.Login,
			Password: string(hashedPassword),
		},
	)
	if err != nil {
		return fmt.Errorf("can't create user")
	}
	return nil
}

// Логин пользователя, проверка существования пользователя
func (uc *UserUsecase) Login(username, password string) (entities.User, error) {
	user, err := uc.Repository.Get(username)
	if err != nil {
		return entities.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return entities.User{}, err
	}
	return *user, nil
}
