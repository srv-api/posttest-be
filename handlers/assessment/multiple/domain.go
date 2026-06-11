package multiple

import (
	s "posttest-be/services/assessment/multiple"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	GetPicture(c echo.Context) error
	Get(c echo.Context) error
}

type domainHandler struct {
	serviceMultiple s.MultipleService
}

func NewMultipleHandler(service s.MultipleService) DomainHandler {
	return &domainHandler{
		serviceMultiple: service,
	}
}
