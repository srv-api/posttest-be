package multiple

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	dto "posttest-be/dto/assessment/multiple"
	"strings"
	"time"

	util "github.com/srv-api/util/s"
)

const (
	maxUploadSize = 10 << 20 // 10MB per file
	uploadDir     = "uploads/multiple"
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
				Index:        i, // ← Pastikan index tersimpan
				Success:      false,
				QuestionText: req.QuestionText,
				Error:        "question_text is required",
			})
			continue
		}

		if len(req.AnswerOptions) == 0 {
			results = append(results, dto.MultipleBatchItemResponse{
				Index:        i, // ← Pastikan index tersimpan
				Success:      false,
				QuestionText: req.QuestionText,
				Error:        "answer_options is required",
			})
			continue
		}

		// Generate ID untuk setiap pertanyaan
		req.ID = util.GenerateRandomString()

		// Set nilai default
		if req.QuestionType == "" {
			req.QuestionType = "multiple_choice"
		}

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

	// Tambahkan success results dengan index yang benar
	// Kita perlu mapping success ke index asli
	successIndex := 0
	for i, req := range reqs {
		// Cari apakah req ini berhasil disimpan
		if successIndex < len(successes) && successes[successIndex].QuestionText == req.QuestionText {
			batchResponse.Results = append(batchResponse.Results, dto.MultipleBatchItemResponse{
				Index:        i, // ← Index asli dari request
				Success:      true,
				QuestionID:   successes[successIndex].ID,
				QuestionText: successes[successIndex].QuestionText,
				Data:         &successes[successIndex],
			})
			successIndex++
		}
	}

	// Tambahkan error results dari repository
	for _, err := range errorsRepo {
		batchResponse.Results = append(batchResponse.Results, dto.MultipleBatchItemResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	// Tambahkan validation errors (sudah punya index)
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
		DetailID:        question.DetailID,
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
			DetailID:        q.DetailID,
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
			DetailID:        q.DetailID,
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
	return s.Repo.Update(id, req)
}

func (s *multipleService) Delete(id string) error {
	return s.Repo.Delete(id)
}

// UploadFile upload file ke server
func UploadFile(file *multipart.FileHeader, questionID string) (string, error) {
	// Buat direktori jika belum ada
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Generate unique filename
	timestamp := time.Now().UnixNano()
	ext := strings.ToLower(filepath.Ext(file.Filename))
	filename := fmt.Sprintf("%s_%d%s", questionID, timestamp, ext)
	filePath := filepath.Join(uploadDir, filename)

	// Buka source file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	// Buat destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	// Copy file
	if _, err = io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	imageURL := "/" + filePath
	return imageURL, nil
}

// ValidateImageFile memvalidasi file gambar
func ValidateImageFile(file *multipart.FileHeader) error {
	// Cek ekstensi file
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExt := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
	}

	if !allowedExt[ext] {
		return fmt.Errorf("file type not allowed. Allowed: jpg, jpeg, png, gif, webp")
	}

	// Cek ukuran file
	if file.Size > maxUploadSize {
		return fmt.Errorf("file too large. Max size: 10MB")
	}

	// Buka file untuk validasi konten
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open file")
	}
	defer src.Close()

	// Baca header file untuk deteksi tipe MIME
	header := make([]byte, 512)
	_, err = src.Read(header)
	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to read file header")
	}

	return nil
}
