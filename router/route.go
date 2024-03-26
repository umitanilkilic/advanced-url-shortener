package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/umitanilkilic/advanced-url-shortener/handler"
	middleware "github.com/umitanilkilic/advanced-url-shortener/middleware"
)

func SetupRoutes(app *fiber.App) {
	// Monitor endpoint
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Advanced Url Shortener"}))

	//app.Get("/:id", redirectUrl) // Redirect to original URL

	// We use the logger middleware here to log all requests
	api := app.Group("/api", logger.New())

	auth := api.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)

	urls := api.Group("/urls", middleware.SecureEndpoint())

	urls.Get("/", handler.GetUrls)
	urls.Post("/", handler.CreateUrl)
	urls.Patch("/:id", handler.UpdateUrl)
	urls.Delete("/:id", handler.DeleteUrl)

	stats := urls.Group("/stats")
	//stats.Get("/:id", getStats)
	_ = stats

}
