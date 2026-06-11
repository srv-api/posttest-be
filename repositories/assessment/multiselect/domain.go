package multiselect

import (
	"gorm.io/gorm"

	dto "posttest-be/dto/assessment/multiselect"
)

type DomainRepository interface {
	Create(req dto.MultiselectRequest) (dto.MultiselectResponse, error)
	GetPicture(req dto.MultiselectRequest) (*dto.MultiselectResponse, error)
	Get(req dto.AccessRoomRequest) ([]dto.MultiselectResponse, error)
}

type multiselectRepository struct {
	DB *gorm.DB
}

func NewMultiselectRepository(DB *gorm.DB) DomainRepository {
	return &multiselectRepository{
		DB: DB,
	}
}
