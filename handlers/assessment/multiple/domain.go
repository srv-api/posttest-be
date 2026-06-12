package multiple

import (
	s "posttest-be/services/assessment/multiple"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	CreateBatch(c echo.Context) error
	GetByID(c echo.Context) error
	GetByDetailID(c echo.Context) error
	Delete(c echo.Context) error
	GetPicture(c echo.Context) error
	Get(c echo.Context) error
	Update(c echo.Context) error
}

type domainHandler struct {
	serviceMultiple s.MultipleService
}

func NewMultipleHandler(service s.MultipleService) DomainHandler {
	return &domainHandler{
		serviceMultiple: service,
	}
}
