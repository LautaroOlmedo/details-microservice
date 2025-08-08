package details

import (
	"context"
	domain "details-microservice/internal/domain/details"
)

//go:generate mockery --name=StorageRepository --output=details --inpackage
type StorageRepository interface {
	GetByID(ctx context.Context, id string) (*domain.Details, error)
}

// If we need to communicate with the product service here would define the interface with their methods to implement.
//	type ProductService interface {
//		GetByID(ctx context.Context, id uint) (*domain.Product, error)
//	}

type Service struct {
	StorageRepository StorageRepository
}

func NewService(storageRepository StorageRepository) *Service {
	return &Service{
		StorageRepository: storageRepository,
	}
}
