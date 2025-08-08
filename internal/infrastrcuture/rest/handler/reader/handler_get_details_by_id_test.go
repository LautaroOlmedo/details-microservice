package reader

import (
	"context"
	domain "details-microservice/internal/domain/details"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var detailsService *MockDetailsService

func TestMain(m *testing.M) {
	detailsService = &MockDetailsService{}
	detailsService.On("GetByID", context.Background(), "550e8400-e29b-41d4-a716-446655440000").Return(&domain.Details{ID: "550e8400-e29b-41d4-a716-446655440000", Description: "Sample description"}, nil)
	code := m.Run()
	os.Exit(code)
}

func TestReaderHandler_HandlerGetDetailsByID(t *testing.T) {

	type testCase struct {
		name               string
		id                 string
		expectedStatusCode int
		expectedBody       string
	}

	testCases := []testCase{
		{
			name:               "PASS. Valid ID",
			id:                 "550e8400-e29b-41d4-a716-446655440000",
			expectedStatusCode: http.StatusOK,
			expectedBody:       `{"ID":"550e8400-e29b-41d4-a716-446655440000","Description":"Sample description"}`,
		},
		{
			name:               "ERROR. Missing ID",
			id:                 "",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "Header details-ID is required",
		},
		{
			name:               "PASS. Not found",
			id:                 "670e8257-a41z-41d4-a716-446655440000",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "Not found",
		},
	}
	readerHandler := NewHandler(detailsService)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/details", nil)
	req.Header.Set("details-ID", "550e8400-e29b-41d4-a716-446655440000")

	rr := httptest.NewRecorder()

	readerHandler.HandlerGetDetailsByID(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var got domain.Details
	err := json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		t.Fatalf("error decoding response: %v", err)
	}

	//if got != expected {
	//	t.Errorf("expected %+v, got %+v", expected, got)
	//}
}

//// Create a test HTTP server that routes requests to our handler.
//handler := &ReaderHandler{
//Details: &MockDetailsService{}, // Use a mock service for testing.
//
//}
//server := httptest.NewServer(handler)
//defer server.Close() // Ensure the server is closed after the test.
//
//req, err := http.NewRequest(http.MethodGet, server.URL, nil)
//if err != nil {
//t.Fatalf("Failed to create request: %v", err)
//}
//
//resp, err := http.DefaultClient.Do(req)
//if err != nil {
//t.Fatalf("Expected no error, got %v", err)
//}
//defer resp.Body.Close() // Close the response body.
//
//if resp.StatusCode != http.StatusOK {
//t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
//}
