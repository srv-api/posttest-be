package multiselect

import (
	dto "posttest-be/dto/assessment/multiselect"
	"posttest-be/entity"
)

func (r *multiselectRepository) GetPicture(req dto.MultiselectRequest) (*dto.MultiselectResponse, error) {
	tr := entity.MultiselectQuestion{
		ImageURL: req.ImageURL,
	}

	if err := r.DB.Where("image_url = ?", tr.ImageURL).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.MultiselectResponse{
		ImageURL: "http://103.150.227.223:2349/" + tr.ImageURL,
	}

	return response, nil
}
