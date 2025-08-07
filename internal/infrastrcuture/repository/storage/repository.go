package storage

import (
	"details-microservice/configuration"
	domain "details-microservice/internal/domain/details"
	"time"
)

type Map struct {
	config     *configuration.Configuration
	detailsMap map[string]domain.Details
}

func NewMapRepository(config *configuration.Configuration) *Map {
	return &Map{
		config: config,
		detailsMap: map[string]domain.Details{
			"550e8400-e29b-41d4-a716-446655440000": {ID: "550e8400-e29b-41d4-a716-446655440000", Description: "Sample description", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"660e8400-e29b-41d4-a716-446655440111": {ID: "660e8400-e29b-41d4-a716-446655440111", Description: "Another description", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}
}

type Slice struct {
}

func NewSliceRepository() *Slice {
	return &Slice{}

}
