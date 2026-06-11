package auth

import (
	dto "posttest-be/dto"

	"github.com/labstack/echo/v4"
	util "github.com/srv-api/util/s"
	res "github.com/srv-api/util/s/response"
)

func (h *domainHandler) Signup(c echo.Context) error {
	var req dto.SignupRequest
	var resp dto.SignupResponse

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	resp, err = h.serviceAuth.Signup(req)
	if err != nil {
		if util.IsDuplicateEntryError(err) {
			return res.ErrorResponse(&res.ErrorConstant.Duplicate).Send(c)
		}
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)
}
