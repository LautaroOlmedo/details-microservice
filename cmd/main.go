package main

import (
	"details-microservice/cmd/dependencies"
	"details-microservice/configuration"
	"details-microservice/internal/infrastrcuture/rest/routes"
	"log"
	"net/http"
)

func main() {

	cfg, err := configuration.Load("../.env")
	if err != nil {
		log.Fatal("Failed to load configuration: " + err.Error())
	}
	dep := dependencies.InitDependencies(cfg)

	// Create a new HTTP server multiplexer
	mux := http.NewServeMux()

	// Set up routes for the application
	routes.SetUpReadRoutes(mux, dep)

	server := &http.Server{
		Addr:    ":8080", // config.SERVER.PORT
		Handler: mux,
	}

	// Start the server
	log.Printf("Server starting on port %s", "8080")
	if err = server.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server: " + err.Error())
	}
}
