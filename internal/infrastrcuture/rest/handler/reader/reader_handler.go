package reader

import (
	"context"
	domain "details-microservice/internal/domain/details"
)

//go:generate mockery --name=DetailsService --output=details --inpackage
type DetailsService interface {
	GetByID(ctx context.Context, id string) (*domain.Details, error)
}

// readerHandler depends on the interfaces, not concrete types.
type ReaderHandler struct {
	Details DetailsService
}

func NewHandler(details DetailsService) *ReaderHandler {
	return &ReaderHandler{
		Details: details,
	}
}
