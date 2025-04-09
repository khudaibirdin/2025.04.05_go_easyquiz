package handlers

import (
	"app/internal/config"
	"app/internal/usecases"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
	UseCase usecases.UserUsecase
	Config  *config.Config
}

func NewUserHandler(uc usecases.UserUsecase, cfg *config.Config) *UserHandler {
	return &UserHandler{
		UseCase: uc,
		Config:  cfg,
	}
}

// Хэндлер Логина
func (h *UserHandler) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	user, err := h.UseCase.Login(req.Login, req.Password)
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
	t, err := jwtToken.SignedString([]byte(h.Config.HTTP.JWTKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(
		fiber.Map{
			"jwt": t,
		},
	)
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req usecases.UserRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}
	err := h.UseCase.Register(req)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.SendStatus(fiber.StatusOK)
}
