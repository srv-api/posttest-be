package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	dto "posttest-be/dto/auth"
	"posttest-be/entity"

	util "github.com/srv-api/util/s"
)

func (s *authService) SignInWithGoogleWeb(req dto.GoogleSignInWebRequest) (*dto.AuthResponse, error) {

	tokenResp, err := http.PostForm(
		"https://oauth2.googleapis.com/token",
		url.Values{
			"code":          {req.Code},
			"client_id":     {os.Getenv("GOOGLE_CLIENT_ID")},
			"client_secret": {os.Getenv("GOOGLE_CLIENT_SECRET")},
			"redirect_uri":  {os.Getenv("GOOGLE_REDIRECT_URI")},
			"grant_type":    {"authorization_code"},
		},
	)
	if err != nil {
		return nil, err
	}
	defer tokenResp.Body.Close()

	var tokenData struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(tokenResp.Body).Decode(&tokenData); err != nil {
		return nil, err
	}
	if tokenData.AccessToken == "" {
		return nil, errors.New("failed to get access token")
	}

	// ===============================
	// 2. GET USER PROFILE
	// ===============================
	reqProfile, _ := http.NewRequest(
		"GET",
		"https://www.googleapis.com/oauth2/v3/userinfo",
		nil,
	)
	reqProfile.Header.Set("Authorization", "Bearer "+tokenData.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(reqProfile)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var profile struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}
	if profile.Email == "" {
		return nil, errors.New("email not found")
	}

	// ===============================
	// 3. USER LOGIC (PUNYA KAMU)
	// ===============================

	encryptedEmail, err := util.Encrypt(profile.Email)
	if err != nil {
		return nil, err
	}

	user, err := s.Repo.FindByEncryptedEmail(encryptedEmail)

	if err != nil && err.Error() == "user not found" {
		// 4. Buat user baru jika belum ada
		if req.Whatsapp == "" {
			return &dto.AuthResponse{
				FullName: profile.Name,
				Email:    encryptedEmail,
				Whatsapp: "",
			}, nil
		}

		secureID, err := generateSecureID()
		if err != nil {
			return nil, err
		}

		encryptedWhatsapp, err := util.Encrypt(req.Whatsapp)
		if err != nil {
			return nil, err
		}

		user = &entity.AccessDoor{
			ID:           secureID,
			Email:        encryptedEmail,
			FullName:     profile.Name,
			Provider:     "google",
			AccessRoleID: "e9Wl2JyVeBM_",
			Whatsapp:     encryptedWhatsapp,
		}

		if err := s.Repo.Create(user); err != nil {
			return nil, err
		}

		user, err = s.Repo.FindByEncryptedEmail(encryptedEmail)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	} else {
		if user.Whatsapp == "/bvTmYgHVZjVt85fktdsXA==" && req.Whatsapp != "/bvTmYgHVZjVt85fktdsXA==" {
			if err := s.Repo.UpdateWhatsapp(user.ID, req.Whatsapp); err != nil {
				return nil, err
			}
			user.Whatsapp = req.Whatsapp
		}
	}

	token, err := s.jwt.GenerateToken(user.ID, user.FullName, user.Whatsapp)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.jwt.GenerateRefreshToken(user.ID, user.FullName, user.Whatsapp)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		ID:            user.ID,
		FullName:      user.FullName,
		Whatsapp:      user.Whatsapp,
		Email:         user.Email,
		Token:         token,
		RefreshToken:  refreshToken,
		TokenVerified: user.Verified.Token,
	}, nil
}
