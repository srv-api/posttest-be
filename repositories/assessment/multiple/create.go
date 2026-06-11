package multiple

import (
	dto "posttest-be/dto/assessment/multiple"
	"posttest-be/entity"
)

type QuestionType string

func (r *multipleRepository) Create(req dto.MultipleRequest) (dto.MultipleResponse, error) {
	create := entity.MultipleQuestion{
		ID:              req.ID,
		QuestionText:    req.QuestionText,
		UserID:          req.UserID,
		DetailID:        req.DetailID,
		CreatedBy:       req.CreatedBy,
		QuestionType:    req.QuestionType,
		AnswerOptions:   req.AnswerOptions,
		Explanation:     req.Explanation,
		PlaceholderText: req.PlaceholderText,
		ImageURL:        req.ImageURL,
	}

	if err := r.DB.Create(&create).Error; err != nil { // Gunakan Create bukan Save
		return dto.MultipleResponse{}, err
	}

	response := dto.MultipleResponse{
		ID:            create.ID,
		UserID:        create.UserID,
		AnswerOptions: create.AnswerOptions,
		DetailID:      create.DetailID,
		CreatedBy:     create.CreatedBy,
		ImageURL:      create.ImageURL,
	}

	return response, nil
}
