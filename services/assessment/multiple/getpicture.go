package multiple

import dto "posttest-be/dto/assessment/multiple"

func (b *multipleService) GetPicture(req dto.MultipleRequest) (*dto.MultipleResponse, error) {
	return b.Repo.GetPicture(req)
}
