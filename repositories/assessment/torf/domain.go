package multiple

import (
	"gorm.io/gorm"

	dto "posttest-be/dto/assessment/multiple"
)

type DomainRepository interface {
	Create(req dto.MultipleRequest) (dto.MultipleResponse, error)
	GetPicture(req dto.MultipleRequest) (*dto.MultipleResponse, error)
	Get(req dto.AccessRoomRequest) ([]dto.MultipleResponse, error)
}

type multipleRepository struct {
	DB *gorm.DB
}

func NewMultipleRepository(DB *gorm.DB) DomainRepository {
	return &multipleRepository{
		DB: DB,
	}
}
