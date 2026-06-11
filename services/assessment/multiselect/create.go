package multiselect

import (
	dto "posttest-be/dto/assessment/multiselect"

	util "github.com/srv-api/util/s"
)

func (s *multiselectService) Create(req dto.MultiselectRequest) (dto.MultiselectResponse, error) {
	create := dto.MultiselectRequest{
		ID:              util.GenerateRandomString(),
		QuestionText:    req.QuestionText,
		UserID:          req.UserID,
		DetailID:        req.DetailID,
		CreatedBy:       req.CreatedBy,
		QuestionType:    req.QuestionType,
		AnswerOptions:   req.AnswerOptions,
		Explanation:     req.Explanation,
		PlaceholderText: req.PlaceholderText,
		ImageURL:        req.ImageURL,
		Image:           req.Image,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.MultiselectResponse{}, err
	}

	response := dto.MultiselectResponse{
		ID:              created.ID,
		UserID:          created.UserID,
		DetailID:        created.DetailID,
		CreatedBy:       created.CreatedBy,
		QuestionType:    created.QuestionType,
		QuestionText:    created.QuestionText,
		AnswerOptions:   created.AnswerOptions,
		Explanation:     created.Explanation,
		PlaceholderText: created.PlaceholderText,
		ImageURL:        created.ImageURL,
	}

	return response, nil
}
