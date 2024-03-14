package urlshortener

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/template/html/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/umitanilkilic/advanced-url-shortener/internal/helper"
	"github.com/umitanilkilic/advanced-url-shortener/internal/model"
)

func RunUrlShortener(cfg map[string]string) error {

	engine := html.New("./web/templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	max, err := strconv.Atoi(cfg["MAX_RATE_LIMIT"])
	if err != nil {
		return fmt.Errorf("error while parsing max_rate_limit: %v", err)
	}
	rateLimitExpiration, err := strconv.Atoi(cfg["RATE_LIMIT_EXPIRATION"])
	if err != nil {
		return fmt.Errorf("error while parsing rate_limit_expiration: %v", err)
	}

	app.Use(limiter.New(limiter.Config{
		Max:        max,
		Expiration: time.Duration(rateLimitExpiration) * time.Second,
		/// SlidingWindow is a rate limiter algorithm that limits the number of requests that can be made in a given time window.
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	app.Post("/shorten", ShortenUrl)
	app.Get("/:shortUrl", GetShortUrl)

	/// Listen on port 8080
	return fmt.Errorf("error while starting web server: %v", app.Listen(cfg["APP_ADDRESS"]+":"+cfg["APP_PORT"]))
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
