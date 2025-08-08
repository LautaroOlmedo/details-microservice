package reader

import (
	"context"
	domain "details-microservice/internal/domain/details"
	service "details-microservice/internal/service/details"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var detailsService *MockDetailsService

func TestMain(m *testing.M) {
	detailsService = &MockDetailsService{}
	detailsService.On("GetByID", context.Background(), "550e8400-e29b-41d4-a716-446655440000").Return(&domain.Details{ID: "550e8400-e29b-41d4-a716-446655440000", Description: "Sample description"}, nil)
	detailsService.On("GetByID", context.Background(), "670e8257-a41z-41d4-a716-446655440000").Return(nil, service.NotFoundError)
	code := m.Run()
	os.Exit(code)
}

func TestReaderHandler_HandlerGetDetailsByID(t *testing.T) {

	type testCase struct {
		name               string
		id                 string
		expectedHeader     string
		expectedStatusCode int
		expectedBody       string
	}

	testCases := []testCase{
		{
			name:               "PASS. Valid ID",
			id:                 "550e8400-e29b-41d4-a716-446655440000",
			expectedHeader:     "details-ID",
			expectedStatusCode: http.StatusOK,
			expectedBody:       `{"ID":"550e8400-e29b-41d4-a716-446655440000","Description":"Sample description"}`,
		},
		{
			name:               "ERROR. Missing details-ID Header",
			id:                 "",
			expectedHeader:     "",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "Header details-ID is required",
		},
		{
			name:               "PASS. Not found",
			id:                 "670e8257-a41z-41d4-a716-446655440000",
			expectedHeader:     "details-ID",
			expectedStatusCode: http.StatusNotFound,
			expectedBody:       "Not found",
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			detailsService.Mock.Test(t)
			readerHandler := NewHandler(detailsService)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/details", nil)
			req.Header.Set(tc.expectedHeader, tc.id)

			rr := httptest.NewRecorder()

			readerHandler.HandlerGetDetailsByID(rr, req)

			resp := rr.Result()
			defer resp.Body.Close()

			if resp.StatusCode != tc.expectedStatusCode {
				t.Fatalf("expected status %d, got %d", tc.expectedStatusCode, resp.StatusCode)
			}
		})
	}

	//var got domain.Details
	//err := json.NewDecoder(resp.Body).Decode(&got)
	//if err != nil {
	//	t.Fatalf("error decoding response: %v", err)
	//}

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
