package rest

import (
	"details-microservice/configuration"
	"details-microservice/internal/infrastrcuture/rest/handler/reader"
	"fmt"
	"github.com/labstack/echo/v4"
)

type RestServer struct {
	config        *configuration.Configuration
	router        *echo.Echo
	readerHandler *reader.Handler
}

func NewRestServer(configuration *configuration.Configuration, readerHandler *reader.Handler) (*RestServer, error) {
	return &RestServer{
		config:        configuration,
		readerHandler: readerHandler,
	}, nil
}

func (s *RestServer) SetUpServer() {
	s.router = echo.New()
	s.SetUpRoutes()
	fmt.Printf("Server started at %s:%s\n", s.config.SERVER.DOMAIN, s.config.SERVER.PORT)
	s.router.Logger.Fatal(s.router.Start(s.config.SERVER.PORT))
}
