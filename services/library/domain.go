package library

import (
	dto "posttest-be/dto"

	"github.com/labstack/echo/v4"
	m "github.com/srv-api/middlewares/middlewares"

	r "posttest-be/repositories/library"
)

type ProductService interface {
	Get(context echo.Context, req *dto.Pagination) dto.Response
}

type productService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewProductService(Repo r.DomainRepository, jwtS m.JWTService) ProductService {
	return &productService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
