package storage

import (
	domain "details-microservice/internal/domain/details"
	"errors"
	"os"
	"testing"
)

var detailsRepo *domain.MockStorageRepository

func TestMain(m *testing.M) {
	detailsRepo = &domain.MockStorageRepository{}
	detailsRepo.On("GetByID", uint64(1)).Return(&domain.Details{ID: 1, Description: "Sample description"}, nil)
	detailsRepo.On("GetByID", uint64(2)).Return(nil, NotFoundError)
	code := m.Run()
	os.Exit(code)
}

func TestMap_GetByID(t *testing.T) {
	type testCase struct {
		name          string
		id            uint64
		expectedError error
	}

	testCases := []testCase{
		{
			name:          "PASS. Valid ID",
			id:            1,
			expectedError: nil,
		},
		{
			name:          "ERROR. Not found ID",
			id:            2,
			expectedError: NotFoundError,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			detailsRepo.Mock.Test(t)
			detailsService := domain.NewService(detailsRepo)

			_, err := detailsService.GetByID(tc.id)
			if !errors.Is(tc.expectedError, err) {
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
