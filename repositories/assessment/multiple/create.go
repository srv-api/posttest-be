package multiple

import (
	"posttest-be/dto/assessment/multiple"
	"posttest-be/entity"
)

func (r *multipleRepository) Create(req multiple.MultipleRequest) (multiple.MultipleResponse, error) {
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

	if err := r.DB.Create(&create).Error; err != nil {
		return multiple.MultipleResponse{}, err
	}

	response := multiple.MultipleResponse{
		ID:              create.ID,
		UserID:          create.UserID,
		DetailID:        create.DetailID,
		QuestionType:    create.QuestionType,
		QuestionText:    create.QuestionText,
		AnswerOptions:   create.AnswerOptions,
		Explanation:     create.Explanation,
		PlaceholderText: create.PlaceholderText,
		CreatedBy:       create.CreatedBy,
		ImageURL:        create.ImageURL,
	}

	return response, nil
}

func (r *multipleRepository) CreateBatch(reqs []multiple.MultipleRequest) ([]multiple.MultipleResponse, []error) {
	var successes []multiple.MultipleResponse
	var errors []error

	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, req := range reqs {
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

		if err := tx.Create(&create).Error; err != nil {
			errors = append(errors, err)
			continue
		}

		response := multiple.MultipleResponse{
			ID:              create.ID,
			UserID:          create.UserID,
			DetailID:        create.DetailID,
			QuestionType:    create.QuestionType,
			QuestionText:    create.QuestionText,
			AnswerOptions:   create.AnswerOptions,
			Explanation:     create.Explanation,
			PlaceholderText: create.PlaceholderText,
			CreatedBy:       create.CreatedBy,
			ImageURL:        create.ImageURL,
		}

		successes = append(successes, response)
	}

	if len(errors) > 0 {
		tx.Rollback()
		return successes, errors
	}

	if err := tx.Commit().Error; err != nil {
		return successes, []error{err}
	}

	return successes, nil
}

func (r *multipleRepository) FindByID(id string) (*entity.MultipleQuestion, error) {
	var question entity.MultipleQuestion
	err := r.DB.Where("id = ?", id).First(&question).Error
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *multipleRepository) FindByUserID(userID string) ([]entity.MultipleQuestion, error) {
	var questions []entity.MultipleQuestion
	err := r.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&questions).Error
	return questions, err
}

func (r *multipleRepository) FindByDetailID(detailID string) ([]entity.MultipleQuestion, error) {
	var questions []entity.MultipleQuestion
	err := r.DB.Where("detail_id = ?", detailID).Order("created_at ASC").Find(&questions).Error
	return questions, err
}

func (r *multipleRepository) Update(id string, req multiple.MultipleUpdateRequest) error {
	updates := map[string]interface{}{
		"question_type":    req.QuestionType,
		"question_text":    req.QuestionText,
		"answer_options":   req.AnswerOptions,
		"explanation":      req.Explanation,
		"placeholder_text": req.PlaceholderText,
		"image_url":        req.ImageURL,
	}
	return r.DB.Model(&entity.MultipleQuestion{}).Where("id = ?", id).Updates(updates).Error
}

func (r *multipleRepository) Delete(id string) error {
	return r.DB.Where("id = ?", id).Delete(&entity.MultipleQuestion{}).Error
}
