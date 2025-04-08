package handlers

import (
	"app/internal/usecases"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var (
	JWT_SECRET = "carbon"
)

type UserHandler struct {
	UseCase usecases.UserUsecase
}

func NewUserHandler(uc usecases.UserUsecase) *UserHandler {
	return &UserHandler{
		UseCase: uc,
	}
}

// Хэндлер Логина
func (h *UserHandler) Login(c *fiber.Ctx) error {
	login := c.FormValue("login")
	password := c.FormValue("password")

	user, err := h.UseCase.Login(login, password)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	jwtClaims := jwt.MapClaims{
		"name": user.Login,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}
	jwtToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwtClaims,
	)
	t, err := jwtToken.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(
		fiber.Map{
			"jwt": t,
		},
	)
}
