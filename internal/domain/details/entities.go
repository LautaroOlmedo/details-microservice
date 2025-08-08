package details

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	InvalidDescriptionError = errors.New("invalid description")
)

type Details struct {
	ID          string    `sql:"id"`
	Description string    `sql:"description"`
	CreatedAt   time.Time `sql:"created_at"`
	UpdatedAt   time.Time `sql:"updated_at"`
	// other fields can be added as needed
}

// Factory patter for a possible detail creation implementation
func NewDetails(description string) (Details, error) {
	if description == "" || len(description) > 255 || len(description) < 5 {
		return Details{}, InvalidDescriptionError
	}
	return Details{
		ID:          uuid.New().String(),
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
