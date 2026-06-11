package multiple

import (
	dto "posttest-be/dto/assessment/multiple"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) Get(c echo.Context) error {
	var req dto.AccessRoomRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(400, "Invalid request")
	}

	medsos, err := h.serviceMultiple.Get(req)
	if err != nil {
		return echo.NewHTTPError(500, "Failed to get medsos")
	}

	return c.JSON(200, echo.Map{
		"status": "success",
		"data":   medsos,
	})
}
