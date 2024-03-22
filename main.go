package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	jwtware "github.com/gofiber/jwt/v3"
)

/* var rd *redis_client.RedisClient
var pg *pg_client.PostgresClient

var ctx = context.Background() */

func main() {
	app := fiber.New()
	jwt := middleware.New(jwtware.New(jwtware.Config{SigningKey: []byte(secret)}))

	//app.Listen(":3000")
}
