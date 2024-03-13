package helper

import (
	"github.com/google/uuid"
)

// you can use this function to generate a short url
func CreateShortUrlID(longUrl *string) int {
	return int(uuid.New().ID())
}
