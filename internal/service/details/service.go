package details

import (
	"context"
	domain "details-microservice/internal/domain/details"
)

//go:generate mockery --name=StorageRepository --output=details --inpackage
type StorageRepository interface {
	GetByID(ctx context.Context, id string) (*domain.Details, error)
}

//	type ProductService interface {
//		GetByID(ctx context.Context, id uint) (*domain.Details, error)
//	}

type Service struct {
	StorageRepository StorageRepository
}

func NewService(storageRepository StorageRepository) *Service {
	return &Service{
		StorageRepository: storageRepository,
	}
}
