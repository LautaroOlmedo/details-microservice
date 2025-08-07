package reader

import (
	"fmt"
	"net/http"
	"regexp"
)

var (
	getOneRe = regexp.MustCompile(`^([a-fA-F0-9-]{36})$`)
)

func (handler *ReaderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request:", r.Method, r.URL.Path)
	switch {
	case r.Method == http.MethodGet && getOneRe.MatchString(r.Header.Get("details-ID")):
		handler.HandlerGetDetailsByID(w, r)
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}
