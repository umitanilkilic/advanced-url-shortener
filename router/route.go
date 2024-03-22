package router

import "github.com/gofiber/fiber"

func SetupRoutes(app *fiber.App) {
	//app.Get("/:id", redirectUrl) // Redirect to original URL

	api := app.Group("/api")

	urls := api.Group("/urls")
	//urls.Get("/", getUrls)
	//urls.Post("/", createUrl)
	//urls.Patch("/:id", updateUrl)
	//urls.Delete("/:id", deleteUrl)
	stats := urls.Group("/stats")
	//stats.Get("/:id", getStats)
	_ = stats

}
