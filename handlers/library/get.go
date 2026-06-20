package library

import (
	"posttest-be/helpers"

	"github.com/labstack/echo/v4"
	res "github.com/srv-api/util/s/response"
)

func (b *domainHandler) Get(c echo.Context) error {
	paginationDTO := helpers.GeneratePaginationRequest(c)

	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	paginationDTO.UserID = userid

	if err := c.Bind(&paginationDTO); err != nil {
		return c.JSON(400, "Invalid request")
	}

	users := b.serviceLibrary.Get(c, paginationDTO)

	return c.JSON(200, users)
}
