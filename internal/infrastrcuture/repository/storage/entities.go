package storage

import (
	domain "details-microservice/internal/domain/details"
	"time"
)

type Details struct {
	ID          uint64    `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var mapDetails = map[uint64]Details{
	1: {ID: 1, Description: "Sample description", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	2: {ID: 2, Description: "Another description", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func fromDomain() {}

func ToDomain(d Details) *domain.Details {
	return &domain.Details{
		ID:          d.ID,
		Description: d.Description,
	}
}
