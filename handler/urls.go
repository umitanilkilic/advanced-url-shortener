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
	userId := getUserID(c)
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

// / Example request: {"name": "example", "long": "https://example.com", "alias": "example", "expires_at": "2022-01-01T00:00:00Z"}
func CreateUrl(c *fiber.Ctx) error {
	type request struct {
		Name      string    `json:"name"`
		Long      string    `json:"long"`
		Alias     string    `json:"alias"`
		ExpiresAt time.Time `json:"expires_at"`
	}

	var input request

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

	if !validateTime(input.ExpiresAt) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
			"errors":  "Expires at is not valid",
		})
	}

	var url model.ShortURL = model.ShortURL{
		Name:      input.Name,
		Long:      input.Long,
		Alias:     input.Alias,
		ExpiresAt: input.ExpiresAt,
	}

	/// Get user id from jwt token
	userId := getUserID(c)
	url.CreatedBy = uint(userId)

	/// Generate url id
	var err error
	url.UrlID, err = helper.GenerateUrlID()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error while generating url id",
			"errors":  err.Error(),
		})
	}

	/// Check if alias is empty
	if input.Alias == "" {
		url.Alias = strconv.Itoa(int(url.UrlID)) /// TODO: Change this to generated alias
	}

	/// Add url to database
	if err := database.DB.Create(&url).Error; err != nil {
		if err.Error() == "Error 1062: Duplicate entry" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid request",
				"errors":  "Alias already exists",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Database error",
			"errors":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Url created",
		"data":    url,
	})
}

func UpdateUrl(c *fiber.Ctx) error {
	type request struct {
		Active    bool      `json:"active"`
		Alias     string    `json:"alias"`
		Name      string    `json:"name"`
		Long      string    `json:"long"`
		ExpiresAt time.Time `json:"expires_at"`
	}
	var input request
	input.Active = true
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request", "errors": err.Error()})
	}

	urlIdentifier := c.Params("id")
	if urlIdentifier == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
			"errors":  "Url id are required",
		})
	}

	url, statusCode := getShortURLByAliasOrId(urlIdentifier)
	if url == (model.ShortURL{}) {
		return c.Status(fiber.StatusNotFound).JSON(statusCode)
	}

	userId := getUserID(c)

	if url.CreatedBy != uint(userId) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
			"errors":  "You are not the owner of this URL",
		})
	}

	/// FIXME: Fix this
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

	if !validateTime(input.ExpiresAt) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
			"errors":  "Expires at is not valid",
		})
	}

	if !validateAlias(input.Alias) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
			"errors":  "Alias already exists or invalid",
		})
	}

	if input.Alias != "" {
		url.Alias = input.Alias
	}
	if input.Name != "" {
		url.Name = input.Name
	}
	if input.Long != "" {
		url.Long = input.Long
	}
	if input.ExpiresAt != (time.Time{}) {
		url.ExpiresAt = input.ExpiresAt
	}

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
		"data":    url,
	})
}

func DeleteUrl(c *fiber.Ctx) error {
	urlIdentifier := c.Params("id")
	if urlIdentifier == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
			"errors":  "Url id are required",
		})
	}

	url, statusCode := getShortURLByAliasOrId(urlIdentifier)
	if url == (model.ShortURL{}) {
		return c.Status(fiber.StatusNotFound).JSON(statusCode)
	}

	userId := getUserID(c)

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

func BulkDeleteUrls(c *fiber.Ctx) error {
	type request struct {
		UrlIds []string `json:"urls_ids"`
	}
	/// Example json request: {"urls_ids": ["url_id1", "url_id2"]}
	var input request
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
			"errors":  err.Error(),
		})
	}
	userId := getUserID(c)

	for _, urlId := range input.UrlIds {
		var url model.ShortURL

		if err := database.DB.Where("url_id = ? AND created_by = ?", urlId, userId).First(&url).Error; err != nil {
			if err.Error() == "record not found" {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"status":  "error",
					"message": "Not found",
					"errors":  "Url not found or you are not the owner of this URL",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Database error",
				"errors":  err.Error(),
			})
		}

		if err := database.DB.Delete(&url).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Database error",
				"errors":  err.Error(),
			})
		}
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Urls deleted",
	})
}

func getUserID(c *fiber.Ctx) (userId float64) {
	userId = c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(float64)
	return
}

func validateURL(originalURL string) bool {
	_, err := url.Parse(originalURL)
	return err == nil
}

func validateName(name string) bool {
	return len(name) >= 3
}

func validateTime(expireTime time.Time) bool {
	return expireTime.After(time.Now())
}

func validateAlias(alias string) (valid bool) {
	response := database.DB.Where("alias = ?", alias).First(&model.ShortURL{})
	if response.RowsAffected == 0 {
		valid = true
	}

	return
}

func getShortURLByAliasOrId(urlIdentifier string) (model.ShortURL, fiber.Map) {
	var url model.ShortURL
	// Check if the urlIdentifier is an integer or not
	if !isAlias(&urlIdentifier) {
		if database.DB.Where("url_id = ?", urlIdentifier).First(&url).Error != nil {
			return model.ShortURL{}, fiber.Map{
				"status":  "error",
				"message": "Not found",
				"errors":  "Url not found",
			}
		}
		return url, fiber.Map{}
	}
	if database.DB.Where("alias = ?", urlIdentifier).First(&url).Error != nil {
		return model.ShortURL{}, fiber.Map{
			"status":  "error",
			"message": "Not found",
			"errors":  "Url not found",
		}
	}

	return url, nil
}

func isAlias(identifier *string) bool {
	if _, err := strconv.Atoi(*identifier); err != nil {
		return true
	}
	return false
}
