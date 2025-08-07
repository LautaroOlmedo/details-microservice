package storage

import (
	"context"
	domain "details-microservice/internal/domain/details"
)

func (s *Slice) GetByID(ctx context.Context, id string) (*domain.Details, error) {
	for _, detail := range s.detailsSlice {
		if detail.ID == id {
			return &detail, nil
		}
	}
	return nil, nil
}
