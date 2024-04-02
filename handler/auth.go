package handler

import (
	"errors"
	"net/mail"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/umitanilkilic/advanced-url-shortener/config"
	"github.com/umitanilkilic/advanced-url-shortener/database"
	"github.com/umitanilkilic/advanced-url-shortener/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var errUserNotFound error = errors.New("user not found")

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/// TODO: Email & password validation should be done in the frontend
/// TODO: Email verification system should be implemented
/// TODO: Password reset system should be implemented

/// Register function is simply creating a new user in the database

func Register(c *fiber.Ctx) error {
	var input model.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request", "errors": err.Error()})
	}
	/// Validate the input
	if isValidEmail(input.Email) || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Email and password are required"})
	}

	// Check if user exists
	_, err := findUserByEmail(input.Email)
	if errors.Is(err, errUserNotFound) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "User already exists"})
	} else if err != nil && !errors.Is(err, errUserNotFound) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Database error"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Could not hash the password"})
	}

	input.Password = string(hashedPassword)

	// Create user
	database.DB.Create(&input)

	return c.JSON(fiber.Map{"status": "success", "message": "User created", "data": input})
}

func Login(c *fiber.Ctx) error {

	var input UserCredentials

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request", "errors": err.Error()})
	}

	/// Validate the input
	if !isValidEmail(input.Email) || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid password or email"})
	}

	userEmail := input.Email
	pass := input.Password

	// Check if user exists
	user, err := findUserByEmail(userEmail)
	//Check user credentials
	if err != nil || !verifyPasswordHash(pass, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid credentials"})
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	/// Part of the token which is the payload
	claims := token.Claims.(jwt.MapClaims)

	// Set token claims
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["sub"] = user.ID

	// Sign the token with a authenticationKey
	//TODO: That way of getting the secret key is not be the best way to do it in a production environment
	authenticationKey := (*config.Config)["AUTH_SECRET"]
	signedToken, err := token.SignedString([]byte(authenticationKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "success login", "data": signedToken})
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func findUserByEmail(email string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where(&model.User{Email: email}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errUserNotFound
		}
		return nil, errors.New("database error")
	}
	return &user, nil
}

func verifyPasswordHash(userPassword string, passwordHash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(userPassword)); err != nil {
		return false
	}
	return true
}
