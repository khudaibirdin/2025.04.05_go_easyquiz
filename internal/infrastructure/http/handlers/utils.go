package handlers

import (
	"time"

	"app/internal/config"
	"app/internal/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func SetSuccessResponse(ctx *fiber.Ctx, status bool, detail interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"success": true,
			"detail":  detail,
		},
	)
}

func SetBadRequestResponse(ctx *fiber.Ctx, status bool, detail interface{}) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(
		fiber.Map{
			"success": true,
			"detail":  detail,
		},
	)
}

type Token struct {
	Key  string
	User *entities.User
}

// Создание токена по ID пользователя
func (t Token) Generate() (string, error) {
	jwtClaims := jwt.MapClaims{
		"id":  t.User.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	jwtToken := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		jwtClaims,
	)
	return jwtToken.SignedString(config.Get().HTTP.Privatekey)
}

// Добавленный Middleware для внутреннего определения ID пользователя
// Просаживается в fiber.ctx.Locals("id")
func JWTMiddleware(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}
	id, ok := claims["id"].(float64)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid user ID in token",
		})
	}
	ctx.Locals("userID", uint(id))
	return ctx.Next()
}
