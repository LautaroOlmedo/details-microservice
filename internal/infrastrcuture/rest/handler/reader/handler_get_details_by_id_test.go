package reader

import (
	domain "details-microservice/internal/domain/details"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReaderHandler_HandlerGetDetailsByID(t *testing.T) {
	expected := domain.Details{
		ID:          "123",
		Description: "Sample Description",
	}

	handler := &ReaderHandler{
		Details: &MockDetailsService{},
	}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/details", nil)
	req.Header.Set("details-ID", "123")

	rr := httptest.NewRecorder()

	handler.HandlerGetDetailsByID(rr, req)

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

	if got != expected {
		t.Errorf("expected %+v, got %+v", expected, got)
	}
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
