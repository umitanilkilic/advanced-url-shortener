package middleware

import (
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt/v5"
)

func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}
