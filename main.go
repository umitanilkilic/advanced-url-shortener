package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umitanilkilic/advanced-url-shortener/database"
	"github.com/umitanilkilic/advanced-url-shortener/router"
)

func main() {
	app := fiber.New()

	err := database.SetupDatabaseConnection()
	if err != nil {
		panic(err)
	}
	router.SetupRoutes(app)

	app.Listen(":3000")
}
