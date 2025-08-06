package details

type GetByIDInput struct {
	id uint64
}

type GetByIDOutput struct {
	*Details
}

func (s *Service) GetByID(id uint64) (*Details, error) {
	return s.StorageRepository.GetByID(id)
}
