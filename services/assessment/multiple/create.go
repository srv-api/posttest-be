package multiple

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	dto "posttest-be/dto/assessment/multiple"
	"strings"
	"time"

	util "github.com/srv-api/util/s"
)

func (s *multipleService) Create(req dto.MultipleRequest) (dto.MultipleResponse, error) {
	// Proses base64 image jika ada
	var imageURL string
	var err error
	if req.ImageBase64 != "" {
		imageURL, err = saveBase64Image(req.ImageBase64, req.ImageFilename, req.ID)
		if err != nil {
			return dto.MultipleResponse{}, fmt.Errorf("failed to save image: %w", err)
		}
	}

	create := dto.MultipleRequest{
		ID:              util.GenerateRandomString(),
		QuestionText:    req.QuestionText,
		UserID:          req.UserID,
		CreatedBy:       req.CreatedBy,
		QuestionType:    req.QuestionType,
		AnswerOptions:   req.AnswerOptions,
		Explanation:     req.Explanation,
		PlaceholderText: req.PlaceholderText,
		ImageURL:        imageURL,
	}

	if create.QuestionType == "" {
		create.QuestionType = "multiple_choice"
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.MultipleResponse{}, err
	}

	return created, nil
}

func (s *multipleService) CreateBatch(reqs []dto.MultipleRequest) (dto.MultipleBatchResponse, error) {
	if len(reqs) == 0 {
		return dto.MultipleBatchResponse{}, errors.New("no questions provided")
	}

	if len(reqs) > 100 {
		return dto.MultipleBatchResponse{}, errors.New("maximum 100 questions per batch")
	}

	var preparedReqs []dto.MultipleRequest
	var results []dto.MultipleBatchItemResponse

	for i, req := range reqs {
		// Validasi required field
		if req.QuestionText == "" {
			results = append(results, dto.MultipleBatchItemResponse{
				Index:        i,
				Success:      false,
				QuestionText: req.QuestionText,
				Error:        "question_text is required",
			})
			continue
		}

		if len(req.AnswerOptions) == 0 {
			results = append(results, dto.MultipleBatchItemResponse{
				Index:        i,
				Success:      false,
				QuestionText: req.QuestionText,
				Error:        "answer_options is required",
			})
			continue
		}

		// Generate ID untuk setiap pertanyaan
		req.ID = util.GenerateRandomString()

		// Set nilai default jika kosong
		if req.QuestionType == "" {
			req.QuestionType = "multiple_choice"
		}

		// Proses base64 image jika ada
		var imageURL string
		var err error
		if req.ImageBase64 != "" {
			imageURL, err = saveBase64Image(req.ImageBase64, req.ImageFilename, req.ID)
			if err != nil {
				results = append(results, dto.MultipleBatchItemResponse{
					Index:        i,
					Success:      false,
					QuestionText: req.QuestionText,
					Error:        fmt.Sprintf("failed to save image: %v", err),
				})
				continue
			}
		}
		req.ImageURL = imageURL

		preparedReqs = append(preparedReqs, req)
	}

	// Jika semua gagal validasi
	if len(preparedReqs) == 0 {
		return dto.MultipleBatchResponse{
			SuccessCount: 0,
			FailedCount:  len(results),
			TotalCount:   len(reqs),
			Results:      results,
		}, nil
	}

	// Simpan ke database
	successes, errorsRepo := s.Repo.CreateBatch(preparedReqs)

	// Buat response
	batchResponse := dto.MultipleBatchResponse{
		SuccessCount: len(successes),
		FailedCount:  len(errorsRepo) + len(results),
		TotalCount:   len(reqs),
		Results:      []dto.MultipleBatchItemResponse{},
	}

	// Tambahkan success results
	for _, success := range successes {
		batchResponse.Results = append(batchResponse.Results, dto.MultipleBatchItemResponse{
			Success:      true,
			QuestionID:   success.ID,
			QuestionText: success.QuestionText,
			Data:         &success,
		})
	}

	// Tambahkan error results dari repository
	for _, err := range errorsRepo {
		batchResponse.Results = append(batchResponse.Results, dto.MultipleBatchItemResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	// Tambahkan validation errors
	batchResponse.Results = append(batchResponse.Results, results...)

	return batchResponse, nil
}

func (s *multipleService) GetByID(id string) (*dto.MultipleResponse, error) {
	question, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.MultipleResponse{
		ID:              question.ID,
		UserID:          question.UserID,
		CreatedBy:       question.CreatedBy,
		QuestionType:    question.QuestionType,
		QuestionText:    question.QuestionText,
		AnswerOptions:   question.AnswerOptions,
		Explanation:     question.Explanation,
		PlaceholderText: question.PlaceholderText,
		ImageURL:        question.ImageURL,
	}, nil
}

func (s *multipleService) GetByUserID(userID string) ([]dto.MultipleResponse, error) {
	questions, err := s.Repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var responses []dto.MultipleResponse
	for _, q := range questions {
		responses = append(responses, dto.MultipleResponse{
			ID:              q.ID,
			UserID:          q.UserID,
			CreatedBy:       q.CreatedBy,
			QuestionType:    q.QuestionType,
			QuestionText:    q.QuestionText,
			AnswerOptions:   q.AnswerOptions,
			Explanation:     q.Explanation,
			PlaceholderText: q.PlaceholderText,
			ImageURL:        q.ImageURL,
		})
	}
	return responses, nil
}

func (s *multipleService) GetByDetailID(detailID string) ([]dto.MultipleResponse, error) {
	questions, err := s.Repo.FindByDetailID(detailID)
	if err != nil {
		return nil, err
	}

	var responses []dto.MultipleResponse
	for _, q := range questions {
		responses = append(responses, dto.MultipleResponse{
			ID:              q.ID,
			UserID:          q.UserID,
			CreatedBy:       q.CreatedBy,
			QuestionType:    q.QuestionType,
			QuestionText:    q.QuestionText,
			AnswerOptions:   q.AnswerOptions,
			Explanation:     q.Explanation,
			PlaceholderText: q.PlaceholderText,
			ImageURL:        q.ImageURL,
		})
	}
	return responses, nil
}

func (s *multipleService) Update(id string, req dto.MultipleUpdateRequest) error {
	// Proses base64 image jika ada
	if req.ImageBase64 != "" {
		imageURL, err := saveBase64Image(req.ImageBase64, req.ImageFilename, id)
		if err != nil {
			return fmt.Errorf("failed to save image: %w", err)
		}
		req.ImageURL = imageURL
	}
	return s.Repo.Update(id, req)
}

func (s *multipleService) Delete(id string) error {
	return s.Repo.Delete(id)
}

// saveBase64Image menyimpan image base64 ke file
func saveBase64Image(base64Str, filename, questionID string) (string, error) {
	// Remove data:image/png;base64, prefix if exists
	parts := strings.Split(base64Str, ",")
	if len(parts) > 1 {
		base64Str = parts[1]
	}

	// Decode base64
	imageData, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %w", err)
	}

	// Detect MIME type
	mimeType := http.DetectContentType(imageData)
	var ext string
	switch mimeType {
	case "image/jpeg":
		ext = ".jpg"
	case "image/png":
		ext = ".png"
	case "image/gif":
		ext = ".gif"
	case "image/webp":
		ext = ".webp"
	default:
		ext = ".jpg"
	}

	// Buat direktori jika belum ada
	uploadDir := "uploads/multiple"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Generate unique filename
	timestamp := time.Now().UnixNano()
	var finalFilename string
	if filename != "" {
		// Gunakan filename asli tapi tambahkan timestamp dan questionID
		nameWithoutExt := strings.TrimSuffix(filename, filepath.Ext(filename))
		finalFilename = fmt.Sprintf("%s_%s_%d%s", questionID, nameWithoutExt, timestamp, ext)
	} else {
		finalFilename = fmt.Sprintf("%s_%d%s", questionID, timestamp, ext)
	}

	filePath := filepath.Join(uploadDir, finalFilename)

	// Save file
	if err := os.WriteFile(filePath, imageData, 0644); err != nil {
		return "", fmt.Errorf("failed to save image: %w", err)
	}

	imageURL := "/" + filePath
	return imageURL, nil
}
