package auth

type SignupRequest struct {
	ID           string `json:"id"`
	Otp          string `json:"otp" form:"otp" validate:"required,otp"`
	Whatsapp     string `json:"whatsapp" form:"whatsapp" validate:"required,whatsapp"`
	Email        string `json:"email" form:"email" validate:"required,email"`
	Password     string `json:"password" form:"password" validate:"required,password"`
	FullName     string `json:"full_name"`
	Gender       string `json:"gender"`
	BirthDate    string `json:"birthdate"`
	Age          int    `json:"age"`
	Token        string `json:"token"`
	AccessRoleID string `json:"access_role_id"`
}

type SignupResponse struct {
	ID           string `json:"id"`
	FullName     string `json:"full_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
	Whatsapp     string `json:"whatsapp" form:"whatsapp" validate:"required,whatsapp"`
	Email        string `json:"email" form:"email" validate:"required,email"`
	Password     string `json:"-" form:"password" validate:"required,password"`
	Token        string `json:"token"`
	AccessRoleID string `json:"access_role_id"`
}
