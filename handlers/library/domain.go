package library

import (
	s "posttest-be/services/library"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Get(c echo.Context) error
}

type domainHandler struct {
	serviceLibrary s.LibraryService
}

func NewLibraryHandler(service s.LibraryService) DomainHandler {
	return &domainHandler{
		serviceLibrary: service,
	}
}
