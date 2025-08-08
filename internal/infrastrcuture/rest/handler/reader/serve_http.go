package reader

import (
	"fmt"
	"net/http"
)

func (handler *ReaderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request:", r.Method, r.URL.Path)
	switch {
	case r.Method == http.MethodGet:
		handler.HandlerGetDetailsByID(w, r)
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}
