package multiple

import (
	"gorm.io/gorm"

	"posttest-be/dto/assessment/multiple"
	dto "posttest-be/dto/assessment/multiple"
	"posttest-be/entity"
)

type DomainRepository interface {
	Create(req dto.MultipleRequest) (dto.MultipleResponse, error)
	GetPicture(req dto.MultipleRequest) (*dto.MultipleResponse, error)
	Get(req dto.AccessRoomRequest) ([]dto.MultipleResponse, error)
	CreateBatch(reqs []multiple.MultipleRequest) ([]multiple.MultipleResponse, []error)
	FindByID(id string) (*entity.MultipleQuestion, error)
	FindByUserID(userID string) ([]entity.MultipleQuestion, error)
	FindByDetailID(detailID string) ([]entity.MultipleQuestion, error)
	Update(id string, req multiple.MultipleUpdateRequest) error
	Delete(id string) error
}

type multipleRepository struct {
	DB *gorm.DB
}

func NewMultipleRepository(DB *gorm.DB) DomainRepository {
	return &multipleRepository{
		DB: DB,
	}
}
