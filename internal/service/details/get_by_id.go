package details

import (
	"context"
	domain "details-microservice/internal/domain/details"
	"errors"
)

var (
	NotFoundError = errors.New("details not found")
)

func (s *Service) GetByID(ctx context.Context, id string) (*domain.Details, error) {
	details, err := s.StorageRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if details == nil {
		return nil, NotFoundError
	}
	return details, nil
}
