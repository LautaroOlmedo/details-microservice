package details

import (
	"errors"
	"testing"
)

//
//func TestMain(m *testing.M) {
//	code := m.Run()
//	os.Exit(code)
//
//}

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
			description:   "Sample description",
			expectedError: nil,
		},
		{
			name:          "ERROR. Empty description",
			id:            1,
			description:   "",
			expectedError: InvalidDescriptionError,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewDetails(tc.description)

			if !errors.Is(tc.expectedError, err) {
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
