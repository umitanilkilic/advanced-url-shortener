package helper

import (
	"fmt"

	"github.com/google/uuid"
)

// you can use this function to generate a short url
func ShortenUrl(longUrl *string) (string, error) {
	shortUrl := generateId()
	return shortUrl, nil
}

func generateId() string {
	return fmt.Sprint(uuid.New().ID())
}
