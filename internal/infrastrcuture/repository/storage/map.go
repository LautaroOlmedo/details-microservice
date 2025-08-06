package storage

import (
	domain "details-microservice/internal/domain/details"
	"errors"
)

var (
	NotFoundError = errors.New("details not found")
)

func (m *Map) GetByID(id uint64) (*domain.Details, error) {
	detail, exists := m.detailsMap[id]
	if exists {
		return ToDomain(detail), nil
	}
	return nil, NotFoundError
}
