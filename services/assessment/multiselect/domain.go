package multiselect

import (
	dto "posttest-be/dto/assessment/multiselect"
	r "posttest-be/repositories/assessment/multiselect"

	m "github.com/srv-api/middlewares/middlewares"
)

type MultiselectService interface {
	Create(req dto.MultiselectRequest) (dto.MultiselectResponse, error)
	GetPicture(req dto.MultiselectRequest) (*dto.MultiselectResponse, error)
	Get(req dto.AccessRoomRequest) ([]dto.MultiselectResponse, error)
}

type multiselectService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewMultiselectService(Repo r.DomainRepository, jwtS m.JWTService) MultiselectService {
	return &multiselectService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
