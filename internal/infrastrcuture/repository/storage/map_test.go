package storage

import (
	"context"
	domain "details-microservice/internal/domain/details"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var testRepo *Map

func TestMain(m *testing.M) {
	testRepo = &Map{
		detailsMap: map[string]domain.Details{
			"550e8400-e29b-41d4-a716-446655440000": {
				ID: "550e8400-e29b-41d4-a716-446655440000", Description: "Sample description", CreatedAt: time.Now(), UpdatedAt: time.Now(),
			},
		},
	}
	code := m.Run()
	os.Exit(code)
}

func TestMap_GetByID(t *testing.T) {
	t.Run("PASS - Should find detail by ID", func(t *testing.T) {
		t.Parallel()
		expectedID := "550e8400-e29b-41d4-a716-446655440000"
		detail, err := testRepo.GetByID(context.Background(), expectedID)

		assert.NoError(t, err)
		assert.NotNil(t, detail)
		assert.Equal(t, expectedID, detail.ID)
	})

	t.Run("PASS - Should return nil for non-existing ID", func(t *testing.T) {
		t.Parallel()
		nonExistentID := "non-existent-id"
		detail, err := testRepo.GetByID(context.Background(), nonExistentID)
		assert.NoError(t, err)
		assert.Nil(t, detail)
	})

	t.Run("PASS - Should return nil for empty ID", func(t *testing.T) {
		t.Parallel()
		emptyID := ""
		detail, err := testRepo.GetByID(context.Background(), emptyID)
		assert.NoError(t, err)
		assert.Nil(t, detail)
	})
}
