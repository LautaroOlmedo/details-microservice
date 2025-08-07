package details

import (
	"context"
	domain "details-microservice/internal/domain/details"
	"errors"
	"os"
	"testing"
)

var detailsRepo *MockStorageRepository

func TestMain(m *testing.M) {
	detailsRepo = &MockStorageRepository{}
	detailsRepo.On("GetByID", context.Background(), "550e8400-e29b-41d4-a716-446655440000").Return(&domain.Details{ID: "550e8400-e29b-41d4-a716-446655440000", Description: "Sample description"}, nil)
	detailsRepo.On("GetByID", context.Background(), "").Return(nil, nil)
	code := m.Run()
	os.Exit(code)
}

func TestService_GetByID(t *testing.T) {
	type testCase struct {
		name          string
		id            string
		expectedError error
	}

	testCases := []testCase{
		{
			name:          "PASS. Valid ID",
			id:            "550e8400-e29b-41d4-a716-446655440000",
			expectedError: nil,
		},
		{
			name:          "ERROR. Not found details",
			id:            "",
			expectedError: NotFoundError,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			detailsRepo.Mock.Test(t)
			detailsService := NewService(detailsRepo)

			_, err := detailsService.GetByID(context.Background(), tc.id)
			if !errors.Is(tc.expectedError, err) {
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
