package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/umitanilkilic/advanced-url-shortener/config"
	"github.com/umitanilkilic/advanced-url-shortener/database"
	"github.com/umitanilkilic/advanced-url-shortener/router"
)

func main() {

	/// Create a new Fiber instance
	app := fiber.New()
	router.RegisterRoutes(app)

	/// Setup the database connection
	err := database.SetupDatabaseConnection()
	if err != nil {
		panic(err)
	}

	/// Setup the cache connection
	//err = cache.InitializeRedisClient()
	//if err != nil {
	//	panic(err)
	//}

	/// Start the server
	appAddress := (*config.Config)["APP_ADDRESS"]
	log.Fatal(app.Listen(appAddress))
}
