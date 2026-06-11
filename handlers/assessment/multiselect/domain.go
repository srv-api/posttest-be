package multiselect

import (
	s "posttest-be/services/assessment/multiselect"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	GetPicture(c echo.Context) error
	Get(c echo.Context) error
}

type domainHandler struct {
	serviceMultiselect s.MultiselectService
}

func NewMultiselectHandler(service s.MultiselectService) DomainHandler {
	return &domainHandler{
		serviceMultiselect: service,
	}
}
