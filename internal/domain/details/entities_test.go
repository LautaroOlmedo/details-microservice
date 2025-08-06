package details

import (
	"errors"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)

}

func TestNewDetails(t *testing.T) {
	type testCase struct {
		name          string
		id            uint64
		description   string
		expectedError error
	}

	testCases := []testCase{
		{
			name:          "PASS. Valid values",
			id:            1,
			description:   "Sample description",
			expectedError: nil,
		},
		{
			name:          "ERROR. Invalid ID",
			id:            0,
			description:   "Sample description",
			expectedError: InvalidIDError,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewDetails(tc.id, tc.description)

			if !errors.Is(tc.expectedError, err) {
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
