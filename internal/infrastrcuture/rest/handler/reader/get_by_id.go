package reader

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *ReaderHandler) HandlerGetDetailsByID(w http.ResponseWriter, r *http.Request) {
	detailsID := r.Header.Get("details-ID")
	if detailsID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Header details-ID is required"))
		if err != nil {
			return
		}
		return
	}

	//path := strings.TrimPrefix(r.URL.Path, "/")
	//parts := strings.Split(path, "/")
	//id := parts[len(parts)-1]

	fmt.Println("ID from request:", detailsID)

	details, err := h.Details.GetByID(r.Context(), detailsID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(fmt.Sprintf("error getting deatils: %s", err)))
		if err != nil {
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	detailsResponse, err := json.Marshal(details)
	if err != nil {
		_, err = w.Write([]byte(err.Error()))
	}

	_, err = w.Write(detailsResponse)
	if err != nil {
		return
	}
}
