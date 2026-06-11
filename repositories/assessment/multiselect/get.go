package multiselect

import (
	dto "posttest-be/dto/assessment/multiselect"
)

func (r *multiselectRepository) Get(req dto.AccessRoomRequest) ([]dto.MultiselectResponse, error) {
	var multiselectResponses []dto.MultiselectResponse

	return multiselectResponses, nil
}
