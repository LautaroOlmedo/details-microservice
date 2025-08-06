package reader

import (
	domain "details-microservice/internal/domain/details"
	"github.com/labstack/echo/v4"
)

type DetailsService interface {
	GetByID(id uint64) (*domain.Details, error)
}
type Handler struct {
	DetailsService DetailsService
}

func (H *Handler) GetDetailsByID(c echo.Context) error {
	//id := c.Param("id")

	return c.String(500, "Not implemented yet")
}
