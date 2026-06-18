package auth

import (
	dto "posttest-be/dto/auth"
	r "posttest-be/repositories/auth"

	m "github.com/srv-api/middlewares/middlewares"
)

type AuthService interface {
	Signin(req dto.SigninRequest) (*dto.SigninResponse, error)
	SigninByPhoneNumber(req dto.SigninRequest) (*dto.SigninResponse, error)
	SignInWithGoogleWeb(req dto.GoogleSignInWebRequest) (*dto.AuthResponse, error)
	Signup(req dto.SignupRequest) (dto.SignupResponse, error)
	RefreshAccessToken(req dto.RefreshTokenRequest) (string, error)
}

type authService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewAuthService(Repo r.DomainRepository, jwtS m.JWTService) AuthService {
	return &authService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
