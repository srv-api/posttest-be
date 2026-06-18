package auth

import (
	dto "posttest-be/dto/auth"

	"github.com/labstack/echo/v4"
	res "github.com/srv-api/util/s/response"
)

func (u *domainHandler) RefreshToken(c echo.Context) error {
	var req dto.RefreshTokenRequest
	if err := c.Bind(&req); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	req.UserID = userid
	// Validate the refresh token (validate inside the service)
	accessToken, err := u.serviceAuth.RefreshAccessToken(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
	}

	// Prepare the response with the new access token
	resp := dto.RefreshTokenResponse{
		AccessToken: accessToken,
		UserID:      userid,
	}

	// Return success with the new access token
	return res.SuccessResponse(resp).Send(c)
}
