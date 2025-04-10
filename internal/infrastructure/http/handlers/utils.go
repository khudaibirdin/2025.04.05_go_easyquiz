package handlers

import (
	"app/internal/config"
	"app/internal/entities"

	"time"

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
