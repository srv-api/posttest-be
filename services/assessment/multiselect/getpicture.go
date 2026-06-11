package multiselect

import dto "posttest-be/dto/assessment/multiselect"

func (b *multiselectService) GetPicture(req dto.MultiselectRequest) (*dto.MultiselectResponse, error) {
	return b.Repo.GetPicture(req)
}
