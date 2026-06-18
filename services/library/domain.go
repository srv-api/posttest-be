package library

import (
	dto "posttest-be/dto"

	"github.com/labstack/echo/v4"
	m "github.com/srv-api/middlewares/middlewares"

	r "posttest-be/repositories/library"
)

type LibraryService interface {
	Get(context echo.Context, req *dto.Pagination) dto.Response
}

type libraryService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewLibraryService(Repo r.DomainRepository, jwtS m.JWTService) LibraryService {
	return &libraryService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
