package library

import (
	s "posttest-be/services/library"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Get(c echo.Context) error
}

type domainHandler struct {
	serviceProduct s.ProductService
}

func NewProductHandler(service s.ProductService) DomainHandler {
	return &domainHandler{
		serviceProduct: service,
	}
}
