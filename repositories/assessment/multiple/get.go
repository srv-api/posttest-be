package multiple

import (
	dto "posttest-be/dto/assessment/multiple"
)

func (r *multipleRepository) Get(req dto.AccessRoomRequest) ([]dto.MultipleResponse, error) {
	var multipleResponses []dto.MultipleResponse

	return multipleResponses, nil
}
