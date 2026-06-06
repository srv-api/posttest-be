package routes

import (
	"posttest-be/configs"
	h_auth "posttest-be/handlers/auth"
	r_auth "posttest-be/repositories/auth"
	s_auth "posttest-be/services/auth"

	"github.com/labstack/echo/v4"
	"github.com/srv-api/middlewares/middlewares"
)

var (
	DB  = configs.InitDB()
	JWT = middlewares.NewJWTService()

	authR = r_auth.NewAuthRepository(DB)
	authS = s_auth.NewAuthService(authR, JWT)
	authH = h_auth.NewAuthHandler(authS)
)

func New() *echo.Echo {

	e := echo.New()
	e.POST("/auth/web/google", authH.GoogleSignInWeb)

	auth := e.Group("/api", middlewares.ApiKeyMiddleware)
	{
		auth.POST("/signin", authH.Signin)
	}

	return e
}
