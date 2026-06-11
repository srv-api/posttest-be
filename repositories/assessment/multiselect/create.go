package multiselect

import (
	dto "posttest-be/dto/assessment/multiselect"
	"posttest-be/entity"
)

type QuestionType string

func (r *multiselectRepository) Create(req dto.MultiselectRequest) (dto.MultiselectResponse, error) {
	create := entity.MultiselectQuestion{
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
		return dto.MultiselectResponse{}, err
	}

	response := dto.MultiselectResponse{
		ID:            create.ID,
		UserID:        create.UserID,
		AnswerOptions: create.AnswerOptions,
		DetailID:      create.DetailID,
		CreatedBy:     create.CreatedBy,
		ImageURL:      create.ImageURL,
	}

	return response, nil
}
