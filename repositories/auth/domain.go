package auth

import (
	"sync"

	dto "posttest-be/dto"
	"posttest-be/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Signin(req dto.SigninRequest) (*entity.AccessDoor, error)
	UpdateTokenVerified(userID string, otp string, token string) (dto.SigninResponse, error)
	UpdateUser(user *entity.AccessDoor) error
	SigninByPhoneNumber(req dto.SigninRequest) (*entity.AccessDoor, error)
	SaveUser(user *entity.AccessDoor) error
	Create(user *entity.AccessDoor) error
	UpdateWhatsapp(userID string, phone string) error
	FindByEncryptedEmail(encryptedEmail string) (*entity.AccessDoor, error)
	RefreshToken(req dto.RefreshTokenRequest) (*entity.AccessDoor, error)
	Signup(req dto.SignupRequest) (dto.SignupResponse, error)
}

type authRepository struct {
	DB    *gorm.DB
	mu    sync.Mutex
	users map[string]*entity.AccessDoor
}

func NewAuthRepository(DB *gorm.DB) DomainRepository {
	return &authRepository{
		DB: DB,
	}
}
