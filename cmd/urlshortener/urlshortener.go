package urlshortener

import (
	"log"
	"time"

	"github.com/gofiber/template/html/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/umitanilkilic/advanced-url-shortener/internal/config"
	"github.com/umitanilkilic/advanced-url-shortener/internal/helper"
	"github.com/umitanilkilic/advanced-url-shortener/internal/model"
)

func RunUrlShortener(cfg config.FiberConfig) {
	engine := html.New("./web/templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

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
		log.Println("error while parsing request body: ", err)
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	shortUrlID := helper.CreateShortUrlID(&payload.LongUrl)

	url := model.ShortURL{Long: payload.LongUrl, CreatedAt: time.Now().Format("2006-01-02 15:04:05"), ID: shortUrlID}
	err := helper.StoreUrl(url)

	if err != nil {
		log.Println("error while storing url: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"shortUrlId": shortUrlID})
}

func GetShortUrl(c *fiber.Ctx) error {
	shortUrl := c.Params("shortUrl")
	longUrl, err := helper.RetrieveLongUrl(&shortUrl)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	///return c.Redirect(longUrl) Uncomment this line and remove the next line to redirect directly to the long url
	return c.Render("redirect", fiber.Map{"LongURL": longUrl})
}
