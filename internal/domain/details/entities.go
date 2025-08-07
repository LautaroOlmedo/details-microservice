package details

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	InvalidIDError          = errors.New("invalid product ID")
	InvalidDescriptionError = errors.New("invalid description")
)

type Details struct {
	ID          string    `sql:"id"`
	Description string    `sql:"description"`
	CreatedAt   time.Time `sql:"created_at"`
	UpdatedAt   time.Time `sql:"updated_at"`
	// other fields can be added as needed
}

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
