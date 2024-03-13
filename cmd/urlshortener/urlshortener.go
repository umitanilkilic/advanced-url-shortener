package urlshortener

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/umitanilkilic/advanced-url-shortener/internal/config"
	"github.com/umitanilkilic/advanced-url-shortener/internal/helper"
)

func RunUrlShortener(cfg config.FiberConfig) {
	app := fiber.New()

	app.Use(limiter.New(limiter.Config{
		Max:        cfg.RateLimitMax,
		Expiration: time.Duration(cfg.RateLimitExpiration) * time.Second,
		/// SlidingWindow is a rate limiter algorithm that limits the number of requests that can be made in a given time window.
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	app.Post("/shorten", ShortenUrl)
	app.Get("/:shortUrl", GetShortUrl)

	/// Listen on port 8080
	log.Fatal(app.Listen(":" + cfg.AppPort))
}

func ShortenUrl(c *fiber.Ctx) error {
	payload := struct {
		LongUrl string `json:"longUrl"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	shortUrl, err := helper.ShortenUrl(&payload.LongUrl)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"shortUrlId": shortUrl})
}

func GetShortUrl(c *fiber.Ctx) error {
	shortUrl := c.Params("shortUrl")
	longUrl, err := helper.RetrieveLongUrl(&shortUrl)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Redirect(longUrl)
}
