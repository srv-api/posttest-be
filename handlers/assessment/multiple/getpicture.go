package multiple

import (
	"errors"
	dto "posttest-be/dto/assessment/multiple"

	res "github.com/srv-api/util/s/response"

	"github.com/labstack/echo/v4"
)

func (b *domainHandler) GetPicture(c echo.Context) error {
	imagePath := c.Param("*")

	if imagePath == "" {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, errors.New("image path is required")).Send(c)
	}

	// Format path seperti yang ada di database
	dbPath := "/" + imagePath // menjadi "/uploads/medsos/..."

	// Validasi ke database apakah path ini valid
	req := dto.MultipleRequest{
		ImageURL: dbPath,
	}

	_, err := b.serviceMultiple.GetPicture(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, errors.New("image not found in database")).Send(c)
	}

	// Serve file dari local path
	localPath := "./" + imagePath
	return c.File(localPath)
}
