package storage

import (
	"context"
	"testing"
)

func TestSlice_GetByID(t *testing.T) {
	t.Run("PASS - Should find detail by ID", func(t *testing.T) {
		t.Parallel()
		repo := NewSliceRepository(nil)
		expectedID := "550e8400-e29b-41d4-a716-446655440000"
		detail, err := repo.GetByID(context.Background(), expectedID)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if detail == nil || detail.ID != expectedID {
			t.Errorf("Expected detail with ID %s, got %v", expectedID, detail)
		}
	})

}
