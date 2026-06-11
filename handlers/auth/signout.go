package auth

import (
	"net/http"
	"time"

	res "github.com/srv-api/util/s/response"

	"github.com/labstack/echo/v4"
)

func (u *domainHandler) Signout(c echo.Context) error {

	// Hapus access token
	http.SetCookie(c.Response().Writer, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Domain:   ".cashpay.co.id",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		// ❌ HttpOnly false atau dihapus
	})

	// Hapus refresh token
	http.SetCookie(c.Response().Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		Domain:   ".cashpay.co.id",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	return res.SuccessResponse(nil).Send(c)
}
