package multiple

import (
	dto "posttest-be/dto/assessment/multiple"

	util "github.com/srv-api/util/s"
)

func (s *multipleService) Create(req dto.MultipleRequest) (dto.MultipleResponse, error) {
	create := dto.MultipleRequest{
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
		return dto.MultipleResponse{}, err
	}

	response := dto.MultipleResponse{
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
