package helper

import (
	"github.com/google/uuid"
)

func GenerateUrlID() (uint32, error) {
	return uuid.New().ID(), nil
}
