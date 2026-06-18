package auth

import (
	"errors"
	dto "posttest-be/dto/auth"
	"regexp"
	"strings"
	"time"

	res "github.com/srv-api/util/s/response"

	util "github.com/srv-api/util/s"
	"gorm.io/gorm"
)

const MaxLoginAttempts = 3
const SuspendDuration = 5 * time.Minute

func (u *authService) SigninByPhoneNumber(req dto.SigninRequest) (*dto.SigninResponse, error) {
	req.Whatsapp = strings.ToLower(req.Whatsapp)

	req.Whatsapp = FormatWhatsappNumber(req.Whatsapp)

	// Encrypt the Whatsapp
	encryptedWhatsapp, err := util.Encrypt(req.Whatsapp)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	// Set the encrypted Whatsapp in the request
	req.Whatsapp = encryptedWhatsapp

	user, err := u.Repo.SigninByPhoneNumber(req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, res.ErrorBuilder(&res.ErrorConstant.RecordNotFound, err)
		}
		return nil, res.ErrorResponse(err)
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, res.ErrorBuilder(&res.ErrorConstant.RecordNotFound, err)
		}

		return nil, res.ErrorResponse(err)
	}

	if !user.Verified.Verified {

		otp := GenerateRandomNumeric(4)
		tokenVerified := util.GenerateRandomString()

		if err := util.Mail(user.Email, otp); err != nil {
			return nil, err
		}

		if _, err := u.Repo.UpdateTokenVerified(user.ID, otp, tokenVerified); err != nil {
			return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		// Contoh respons:
		response := &dto.SigninResponse{
			VerifiedResp: &dto.AuthUnverifiedResponse{
				TokenVerified: tokenVerified,
				Otp:           otp,
			},
		}
		// // return dto.SigninResponse{
		// // 	TokenVerified: user.Verified.Token,
		// // 	Otp:           user.Verified.Otp,
		// // }, res.ErrorBuilder(&res.ErrorConstant.Unverified, err)
		return response, nil
		// return nil, res.ErrorBuilder(&res.ErrorConstant.Unverified, err)

	}

	// Cek apakah akun disuspend
	if user.Suspended {
		if time.Since(user.LastAttempt) < SuspendDuration {
			return nil, res.ErrorBuilder(&res.ErrorConstant.Suspend, err)
		}
		// Reset status suspend setelah 5 menit
		user.Suspended = false
		user.LoginAttempts = 0
		u.Repo.SaveUser(user)
	}

	// Verifikasi password
	if err != nil || util.VerifyPassword(user.Password, req.Password) != nil {
		user.LoginAttempts++
		user.LastAttempt = time.Now()

		// Jika sudah mencapai batas login gagal, suspend akun
		if user.LoginAttempts >= MaxLoginAttempts {
			user.Suspended = true
			user.LastAttempt = time.Now()
		}

		u.Repo.SaveUser(user)
		return nil, res.ErrorBuilder(&res.ErrorConstant.VerifyPassword, err)
	}

	accesstoken, err := u.jwt.GenerateToken(user.ID, user.FullName, user.Whatsapp)
	refreshtoken, err := u.jwt.GenerateRefreshToken(user.ID, user.FullName, user.Whatsapp)

	return &dto.SigninResponse{
		ID:            user.ID,
		DetailID:      user.Whatsapp,
		FullName:      user.FullName,
		Email:         user.Email,
		AccessToken:   accesstoken,
		RefreshToken:  refreshtoken,
		TokenVerified: user.Verified.Token,
	}, nil
}
func (u *authService) Signin(req dto.SigninRequest) (*dto.SigninResponse, error) {
	req.Email = strings.ToLower(req.Email)

	// Encrypt the email
	encryptedEmail, err := util.Encrypt(req.Email)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	// Set the encrypted email in the request
	req.Email = encryptedEmail

	user, err := u.Repo.Signin(req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, res.ErrorBuilder(&res.ErrorConstant.RecordNotFound, err)
		}
		return nil, res.ErrorResponse(err)
	}

	if !user.Verified.Verified {
		otp := GenerateRandomNumeric(4)
		tokenVerified := util.GenerateRandomString()

		if err := util.Mail(req.Email, otp); err != nil {
			return nil, err
		}

		if _, err := u.Repo.UpdateTokenVerified(user.ID, otp, tokenVerified); err != nil {
			return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		response := &dto.SigninResponse{
			VerifiedResp: &dto.AuthUnverifiedResponse{
				TokenVerified: tokenVerified,
				Otp:           otp,
			},
		}
		return response, nil
	}

	if err != nil || util.VerifyPassword(user.Password, req.Password) != nil {

		// Update the user in the repository
		if err := u.Repo.UpdateUser(user); err != nil {
			return nil, res.ErrorResponse(err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.VerifyPassword, err)
	}

	accesstoken, err := u.jwt.GenerateToken(user.ID, user.FullName, user.Whatsapp)
	refreshtoken, err := u.jwt.GenerateRefreshToken(user.ID, user.FullName, user.Whatsapp)

	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	// Decrypt the email for the response
	decryptedEmail, err := util.Decrypt(user.Email)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &dto.SigninResponse{
		ID:            user.ID,
		DetailID:      user.Whatsapp,
		FullName:      user.FullName,
		Email:         decryptedEmail,
		AccessToken:   accesstoken,
		RefreshToken:  refreshtoken,
		TokenVerified: user.Verified.Token,
	}, nil
}

func FormatWhatsappNumber(phone string) string {
	// Remove non-digit characters
	re := regexp.MustCompile(`\D`)
	phone = re.ReplaceAllString(phone, "")

	// Format the phone number
	if strings.HasPrefix(phone, "0") {
		phone = "+62" + phone[1:]
	} else if strings.HasPrefix(phone, "62") && !strings.HasPrefix(phone, "+") {
		phone = "+" + phone
	} else if !strings.HasPrefix(phone, "+") {
		phone = "+62" + phone
	}

	return phone
}
