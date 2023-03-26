package unique

import (
	"errors"

	"github.com/google/uuid"
)

// ErrInvalidUUID occurs when an UUID is not in a valid form.
var ErrInvalidUUID = errors.New("UUID is not in its proper form")

// NewUUIDv4 returns a new UUID v4.
func NewUUIDv4() string {
	return uuid.NewString()
}

// ValidateUUID validates that the format of an id is valid.
func ValidateUUID(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return ErrInvalidUUID
	}

	return nil
}
