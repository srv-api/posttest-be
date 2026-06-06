package auth

import (
	s "posttest-be/services/auth"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Signin(c echo.Context) error //masuk
	GoogleSignInWeb(c echo.Context) error
}

type domainHandler struct {
	serviceAuth s.AuthService
}

func NewAuthHandler(service s.AuthService) DomainHandler {
	return &domainHandler{
		serviceAuth: service,
	}
}
