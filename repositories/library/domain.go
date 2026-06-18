package library

import (
	"sync"

	dto "posttest-be/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req *dto.Pagination) (RepositoryResult, int)
}

type productRepository struct {
	DB *gorm.DB
	mu sync.Mutex
}

func NewProductRepository(DB *gorm.DB) DomainRepository {
	return &productRepository{
		DB: DB,
	}
}
