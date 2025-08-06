package rest

import "github.com/labstack/echo/v4"

func (s *RestServer) SetUpRoutes() {
	s.MaterialRoutes()
	// Add new routes here, prefered grouped by domain
	//adminGroupe := s.router.Group("/admin")
}

func (s *RestServer) MaterialRoutes() {
	s.router.GET("/details/:id", func(c echo.Context) error {
		return s.readerHandler.GetDetailsByID(c)
	})

}
