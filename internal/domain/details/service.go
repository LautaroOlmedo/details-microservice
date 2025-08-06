package details

//go:generate mockery --name=StorageRepository --output=details --inpackage
type StorageRepository interface {
	GetByID(id uint64) (*Details, error)
}

type ProductService interface {
	ValidateProduct(id uint64) (uint64, error)
}
type Service struct {
	StorageRepository StorageRepository
}

func NewService(storageRepository StorageRepository) *Service {
	return &Service{
		StorageRepository: storageRepository,
	}
}
