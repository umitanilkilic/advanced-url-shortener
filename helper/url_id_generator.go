package helper

import (
	"strconv"

	"github.com/google/uuid"
)

func GenerateUrlID() (string, error) {

	return strconv.Itoa(int(uuid.New().ID())), nil
}
