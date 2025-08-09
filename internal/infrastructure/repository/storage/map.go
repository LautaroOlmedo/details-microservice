package storage

import (
	"context"
	domain "details-microservice/internal/domain/details"
)

func (m *Map) GetByID(ctx context.Context, id string) (*domain.Details, error) {
	detail, exists := m.detailsMap[id]
	if exists {
		return &detail, nil
	}
	return nil, nil
}
