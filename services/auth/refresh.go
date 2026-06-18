package auth

import (
	"errors"

	dto "posttest-be/dto/auth"

	res "github.com/srv-api/util/s/response"
	"gorm.io/gorm"
)

func (u *authService) RefreshAccessToken(req dto.RefreshTokenRequest) (string, error) {
	// Validasi refresh token dan dapatkan user ID
	request := dto.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
		UserID:       req.UserID,
	}

	user, err := u.Repo.RefreshToken(request)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", res.ErrorBuilder(&res.ErrorConstant.RecordNotFound, err)
		}
		return "", res.ErrorResponse(err)
	}

	// Generate token baru
	accessToken, err := u.jwt.GenerateRefreshToken(user.ID, user.FullName, user.Whatsapp)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
