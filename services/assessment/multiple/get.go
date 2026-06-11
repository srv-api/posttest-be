package multiple

import (
	dto "posttest-be/dto/assessment/multiple"
)

func (s *multipleService) Get(req dto.AccessRoomRequest) ([]dto.MultipleResponse, error) {
	multipleResponses, err := s.Repo.Get(req)
	if err != nil {
		return nil, err
	}
	return multipleResponses, nil
}
