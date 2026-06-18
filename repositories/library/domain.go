package library

import (
	"sync"

	dto "posttest-be/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req *dto.Pagination) (RepositoryResult, int)
}

type libraryRepository struct {
	DB *gorm.DB
	mu sync.Mutex
}

func NewLibraryRepository(DB *gorm.DB) DomainRepository {
	return &libraryRepository{
		DB: DB,
	}
}
