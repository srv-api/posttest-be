package auth

type SigninRequest struct {
	Whatsapp string `json:"whatsapp" validate:"required,whatsapp"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" validate:"required"`
	UserID       string `json:"user_id"` // Bisa menambahkan UserID untuk validasi lebih lanjut
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
	UserID      string `json:"user_id"` // Bisa menambahkan UserID untuk validasi lebih lanjut

}

type SigninResponse struct {
	ID            string `json:"id"`
	DetailID      string `json:"detail_id"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	AccessToken   string `json:"access_token"`
	RefreshToken  string `json:"refresh_token"`
	TokenVerified string `json:"token_verified"`
	VerifiedResp  *AuthUnverifiedResponse
	Status        bool `json:"status"`
}

type AuthUnverifiedResponse struct {
	Whatsapp      string `json:"whatsapp"`
	Otp           string `json:"otp"`
	TokenVerified string `json:"token_verified"`
}
