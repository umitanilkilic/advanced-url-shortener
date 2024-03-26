package handler

import (
	"net/url"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/umitanilkilic/advanced-url-shortener/database"
	"github.com/umitanilkilic/advanced-url-shortener/helper"
	"github.com/umitanilkilic/advanced-url-shortener/model"
)

func GetUrls(c *fiber.Ctx) error {
	userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(float64)
	limitQueryParam := c.Query("limit")
	var limit int
	var err error
	if limitQueryParam == "" {
		limit = 10
	} else {
		limit, err = strconv.Atoi(limitQueryParam)
	}

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
			"errors":  "Limit must be a number",
		})
	}

	var urls []model.ShortURL
	if result := database.DB.Where("created_by = ?", userId).Limit(limit).Find(&urls); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Database error",
			"errors":  result.Error.Error,
		})
	}

	return c.JSON(urls)
}

func CreateUrl(c *fiber.Ctx) error {
	var input model.ShortURL

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
			"errors":  err.Error(),
		})
	}

	if !validateName(input.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
			"errors":  "Name are invalid",
		})
	}

	if !validateURL(input.Long) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
			"errors":  "Long Url are not valid",
		})
	}

	userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(float64)
	input.CreatedBy = uint(userId)

	var err error
	input.UrlID, err = helper.GenerateUrlID()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error while generating url id",
			"errors":  err.Error(),
		})
	}

	if err := database.DB.Create(&input).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Database error",
			"errors":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Url created",
		"data":    input,
	})
}

func UpdateUrl(c *fiber.Ctx) error {
	type request struct {
		Active    bool      `json:"active"`
		Name      string    `json:"name"`
		Long      string    `json:"long"`
		ExpiresAt time.Time `json:"expires_at"`
	}
	var input request
	input.Active = true
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request", "errors": err.Error()})
	}

	urlId := c.Params("id")
	if urlId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
			"errors":  "Url id are required",
		})
	}

	var url model.ShortURL

	if result := database.DB.Where("url_id = ?", urlId).First(&url); result.Error != nil {
		if result.Error.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Not found", "errors": "Url not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Database error",
			"errors":  result.Error.Error,
		})
	}

	userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(float64)

	if url.CreatedBy != uint(userId) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
			"errors":  "You are not the owner of this URL",
		})
	}

	if !validateName(input.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request", "errors": "Name are required"})
	}

	if !validateURL(input.Long) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
			"errors":  "Long Url are not valid",
		})
	}

	url.Active = input.Active
	url.Name = input.Name
	url.Long = input.Long
	url.ExpiresAt = input.ExpiresAt

	if database.DB.Save(&url).Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Database error",
			"errors":  "Error while updating url",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Url updated",
		"data":    url})
}

func DeleteUrl(c *fiber.Ctx) error {
	urlId := c.Params("id")
	if urlId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
			"errors":  "Url id are required",
		})
	}

	var url model.ShortURL

	if result := database.DB.Where("url_id = ?", urlId).First(&url); result.Error != nil {
		if result.Error.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Not found", "errors": "Url not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Database error",
			"errors":  result.Error.Error,
		})
	}

	userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(float64)

	if url.CreatedBy != uint(userId) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
			"errors":  "You are not the owner of this URL",
		})
	}

	if database.DB.Delete(&url).Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Database error",
			"errors":  "Error while deleting url",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Url deleted",
	})
}

func validateURL(originalURL string) bool {
	_, err := url.Parse(originalURL)
	return err == nil
}

func validateName(name string) bool {
	return len(name) >= 3
}
