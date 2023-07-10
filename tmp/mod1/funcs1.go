package mod1

import (
	"github.com/google/uuid"
)

func CreateId() string {
	return uuid.New().String()
}
