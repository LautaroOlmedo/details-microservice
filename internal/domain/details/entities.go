package details

import (
	"errors"
)

var (
	InvalidIDError          = errors.New("invalid product ID")
	InvalidDescriptionError = errors.New("invalid description")
)

type Details struct {
	ID          uint64 `json:"id"`
	Description string `json:"description"`
	// other fields can be added as needed
}

func NewDetails(id uint64, description string) (Details, error) {
	if id == 0 {
		return Details{}, InvalidIDError
	}
	if description == "" || len(description) > 255 || len(description) < 5 {
		return Details{}, errors.New("invalid description")

	}
	return Details{
		ID:          id,
		Description: description,
	}, nil
}
