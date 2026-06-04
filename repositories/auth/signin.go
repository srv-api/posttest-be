package auth

import (
	"errors"
	dto "posttest-be/dto"
	"time"

	"posttest-be/entity"

	util "github.com/srv-api/util/s"
)

func (u *authRepository) SigninByPhoneNumber(req dto.SigninRequest) (*entity.AccessDoor, error) {
	var existingUser entity.AccessDoor
	err := u.DB.Preload("Verified").Preload("Merchant").Where("whatsapp = ?", req.Whatsapp).First(&existingUser).Error
	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}

func (u *authRepository) Signin(req dto.SigninRequest) (*entity.AccessDoor, error) {
	var existingUser entity.AccessDoor
	err := u.DB.Preload("Verified").Preload("Merchant").Where("email = ?", req.Email).First(&existingUser).Error
	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}

func (u *authRepository) UpdateUser(user *entity.AccessDoor) error {
	return u.DB.Model(user).Updates(user).Error
}

func (u *authRepository) UpdateTokenVerified(userID string, otp string, token string) (dto.SigninResponse, error) {
	verified := entity.UserVerified{
		ID:        util.GenerateRandomString(),
		UserID:    userID,
		Otp:       otp,
		Token:     token,
		ExpiredAt: time.Now().Add(4 * time.Minute),
	}

	if err := u.DB.Save(&verified).Error; err != nil {
		return dto.SigninResponse{}, err
	}

	response := dto.SigninResponse{
		VerifiedResp: &dto.AuthUnverifiedResponse{
			TokenVerified: verified.Token,
			Otp:           otp,
		},
	}

	return response, nil
}

func (repo *authRepository) SaveUser(user *entity.AccessDoor) error {
	if err := repo.DB.Save(user).Error; err != nil {
		return errors.New("failed to save user")
	}
	return nil
}
