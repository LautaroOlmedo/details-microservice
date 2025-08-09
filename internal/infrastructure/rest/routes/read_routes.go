package routes

import (
	"details-microservice/cmd/dependencies"
	"details-microservice/internal/infrastructure/rest/handler/reader"
	"net/http"
)

func SetUpReadRoutes(mux *http.ServeMux, dep dependencies.Dependencies) {
	readHandler := reader.NewHandler(dep.ReaderHandler.Details)
	mux.HandleFunc("/ping", readHandler.Ping)
	mux.HandleFunc("/api/v1/details/", readHandler.ServeHTTP)
}
