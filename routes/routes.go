package routes

import (
	"posttest-be/configs"
	h_auth "posttest-be/handlers/auth"
	r_auth "posttest-be/repositories/auth"
	s_auth "posttest-be/services/auth"

	h_multiple "posttest-be/handlers/assessment/multiple"
	r_multiple "posttest-be/repositories/assessment/multiple"
	s_multiple "posttest-be/services/assessment/multiple"

	h_multiselect "posttest-be/handlers/assessment/multiselect"
	r_multiselect "posttest-be/repositories/assessment/multiselect"
	s_multiselect "posttest-be/services/assessment/multiselect"

	"github.com/labstack/echo/v4"
	"github.com/srv-api/middlewares/middlewares"
)

var (
	DB  = configs.InitDB()
	JWT = middlewares.NewJWTService()

	authR = r_auth.NewAuthRepository(DB)
	authS = s_auth.NewAuthService(authR, JWT)
	authH = h_auth.NewAuthHandler(authS)

	multipleR = r_multiple.NewMultipleRepository(DB)
	multipleS = s_multiple.NewMultipleService(multipleR, JWT)
	multipleH = h_multiple.NewMultipleHandler(multipleS)

	multiselectR = r_multiselect.NewMultiselectRepository(DB)
	multiselectS = s_multiselect.NewMultiselectService(multiselectR, JWT)
	multiselectH = h_multiselect.NewMultiselectHandler(multiselectS)
)

func New() *echo.Echo {

	e := echo.New()
	e.POST("/d/web/google", authH.GoogleSignInWeb)
	e.GET("/picture/*", multipleH.GetPicture)
	e.GET("/picture/*", multiselectH.GetPicture)
	e.POST("/d/logout", authH.Signout)

	multiple := e.Group("/d", middlewares.AuthorizeJWT(JWT))
	{
		// multiple.POST("/create/multiple", multipleH.Create)
		multiple.GET("/get/multiple", multipleH.Get)
		multiple.POST("/create", multipleH.Create)
		multiple.POST("/batch", multipleH.CreateBatch) // Endpoint batch utama
		multiple.GET("/:id", multipleH.GetByID)
		multiple.GET("/detail/:detail_id", multipleH.GetByDetailID)
		multiple.PUT("/:id", multipleH.Update)
		multiple.DELETE("/:id", multipleH.Delete)

	}

	multiselect := e.Group("/d", middlewares.AuthorizeJWT(JWT))
	{
		multiselect.POST("/create/multiselect", multiselectH.Create)
		multiselect.GET("/get/multiselect", multiselectH.Get)
	}

	auth := e.Group("/d", middlewares.ApiKeyMiddleware)
	{
		auth.POST("/signin", authH.Signin)
		auth.POST("/refresh", authH.RefreshToken)
		auth.POST("/signup", authH.Signup)

	}

	return e
}
