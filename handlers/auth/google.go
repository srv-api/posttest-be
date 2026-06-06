package auth

import (
	"net/http"
	dto "posttest-be/dto"

	"github.com/labstack/echo/v4"
	util "github.com/srv-api/util/s"
	res "github.com/srv-api/util/s/response"
)

func (h *domainHandler) GoogleSignInWeb(c echo.Context) error {
	var req dto.GoogleSignInWebRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	resp, err := h.serviceAuth.SignInWithGoogleWeb(req)
	if err != nil {
		if util.IsDuplicateEntryError(err) {
			return res.ErrorResponse(&res.ErrorConstant.Duplicate).Send(c)
		}
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(resp).Send(c)
}
