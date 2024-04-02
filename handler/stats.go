package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umitanilkilic/advanced-url-shortener/database"
	"github.com/umitanilkilic/advanced-url-shortener/model"
)

func GetStats(c *fiber.Ctx) error {
	urlIdentifier := c.Params("id")
	if urlIdentifier == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
			"errors":  "Url id are required",
		})
	}

	var stats []model.UrlStats
	if database.DB.First("url_id = ?", urlIdentifier).Find(&stats).Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found",
			"errors":  "Url not found",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   stats,
	})
}
