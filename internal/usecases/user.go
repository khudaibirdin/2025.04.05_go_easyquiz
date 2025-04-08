package usecases

import (
	"app/internal/entities"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Add(entities.User) (uint, error)
	Get(username string) (entities.User, error)
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
	_, err := uc.Repository.Get(user.Login)
	if err != nil {
		return fmt.Errorf("user is already exists")
	}
	_, err = uc.Repository.Add(
		entities.User{
			Login:    user.Login,
			Password: user.Password,
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
	return user, nil
}
