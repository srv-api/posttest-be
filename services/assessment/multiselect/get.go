package multiselect

import (
	dto "posttest-be/dto/assessment/multiselect"
)

func (s *multiselectService) Get(req dto.AccessRoomRequest) ([]dto.MultiselectResponse, error) {
	multiselectResponses, err := s.Repo.Get(req)
	if err != nil {
		return nil, err
	}
	return multiselectResponses, nil
}
