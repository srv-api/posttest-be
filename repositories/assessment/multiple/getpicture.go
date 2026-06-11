package multiple

import (
	dto "posttest-be/dto/assessment/multiple"
	"posttest-be/entity"
)

func (r *multipleRepository) GetPicture(req dto.MultipleRequest) (*dto.MultipleResponse, error) {
	tr := entity.MultipleQuestion{
		ImageURL: req.ImageURL,
	}

	if err := r.DB.Where("image_url = ?", tr.ImageURL).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.MultipleResponse{
		ImageURL: "http://103.150.227.223:2349/" + tr.ImageURL,
	}

	return response, nil
}
