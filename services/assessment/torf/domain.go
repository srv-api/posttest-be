package multiple

import (
	dto "posttest-be/dto/assessment/multiple"
	r "posttest-be/repositories/assessment/multiple"

	m "github.com/srv-api/middlewares/middlewares"
)

type MultipleService interface {
	Create(req dto.MultipleRequest) (dto.MultipleResponse, error)
	GetPicture(req dto.MultipleRequest) (*dto.MultipleResponse, error)
	Get(req dto.AccessRoomRequest) ([]dto.MultipleResponse, error)
}

type multipleService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewMultipleService(Repo r.DomainRepository, jwtS m.JWTService) MultipleService {
	return &multipleService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
