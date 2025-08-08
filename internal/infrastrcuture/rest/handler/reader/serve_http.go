package reader

import (
	"net/http"
)

func (handler *ReaderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:
		handler.HandlerGetDetailsByID(w, r)
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}
