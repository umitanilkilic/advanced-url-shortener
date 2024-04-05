package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/umitanilkilic/advanced-url-shortener/database"
	"github.com/umitanilkilic/advanced-url-shortener/model"
)

func RedirectToUrl(c *fiber.Ctx) error {
	urlIdentifier := c.Params("id")
	if urlIdentifier == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid reque1st",
			"errors":  "Url id are required",
		})
	}

	/// TODO: Get the original URL from the cache

	url, statusCode := getShortURLByAliasOrId(urlIdentifier)
	if url == (model.ShortURL{}) {
		return c.Status(fiber.StatusNotFound).JSON(statusCode)
	}

	if !url.Active {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found",
			"errors":  "Url is not active",
		})
	}

	if (url.ExpiresAt != time.Time{} && url.ExpiresAt.Before(time.Now())) {
		url.Active = false
		if database.DB.Save(&url).Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Database error",
				"errors":  "Error while updating url",
			})
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found",
			"errors":  "Url is expired",
		})

	}
	clientIP := getClientIP(c)
	clientUserAgent := getClientUserAgent(c)

	///TODO: Get client ip address, user agent and other information then save it to the database
	activity := model.UrlStats{
		UrlID:        url.UrlID,
		ActivityTime: time.Now(),
		ActivityType: model.ActivityTypeRedirect,
		IpAddress:    clientIP,
		UserAgent:    clientUserAgent,
	}

	if err := database.DB.Create(&activity).Error; err != nil {
		log.Error(err)
	}

	c.SendString("Redirecting to original URL:" + url.Long)

	return c.Redirect(url.Long)

}

func getClientIP(c *fiber.Ctx) string {
	clientIP := c.Get("X-Real-IP")
	if clientIP == "" {
		clientIP = c.IP()
	}
	return clientIP
}

func getClientUserAgent(c *fiber.Ctx) string {
	return c.Get("User-Agent")
}
