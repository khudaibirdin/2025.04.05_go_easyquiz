package usecases

import (
	"app/internal/entities"
	"fmt"
)

type UserRepository interface {
	Add(entities.User) error
	Get(string) (entities.User, bool)
}

type UserUsecase struct {
	Repository UserRepository
}

func NewUserUsecase(r UserRepository) *UserUsecase {
	return &UserUsecase{Repository: r}
}

// func (uc *UserUsecase) Login(login, password string) (string, bool) {
// 	user, userExists := uc.Repository.Get(name)
// 	if userExists != true {
// 		return false
// 	}

// 	return true
// }

type UserRegisterRequest struct {
	Login    string
	Password string
}

func (uc *UserUsecase) Register(user UserRegisterRequest) error {
	_, userExists := uc.Repository.Get(user.Login)
	if userExists {
		return fmt.Errorf("user is already exists")
	}
	err := uc.Repository.Add(
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
