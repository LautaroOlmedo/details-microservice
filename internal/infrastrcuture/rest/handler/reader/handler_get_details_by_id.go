package reader

import (
	service "details-microservice/internal/service/details"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (h *ReaderHandler) HandlerGetDetailsByID(w http.ResponseWriter, r *http.Request) {
	detailsID := r.Header.Get("Details-ID")
	if detailsID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Header Details-ID is required"))
		if err != nil {
			return
		}
		return
	}

	// other possible validation for detailsID can be added here
	//path := strings.TrimPrefix(r.URL.Path, "/")
	//parts := strings.Split(path, "/")
	//id := parts[len(parts)-4]

	details, err := h.Details.GetByID(r.Context(), detailsID)
	if err != nil {
		if errors.Is(err, service.NotFoundError) {
			w.WriteHeader(http.StatusNotFound)
			_, err = w.Write([]byte("Not found"))
			if err != nil {
				return
			}
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(fmt.Sprintf("error getting details: %s", err.Error())))
		if err != nil {
			return
		}
		return
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
