package multiple

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	dto "posttest-be/dto/assessment/multiple"
	"posttest-be/services/assessment/multiple"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) CreateBatch(c echo.Context) error {
	// Parse multipart form dengan max size 100MB untuk semua file
	if err := c.Request().ParseMultipartForm(100 << 20); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Failed to parse form",
			"error":   err.Error(),
		})
	}

	// Ambil questions dari form field
	questionsJSON := c.FormValue("questions")
	if questionsJSON == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Questions field is required",
		})
	}

	// Parse JSON ke struct
	var batchReq dto.MultipleBatchRequest
	if err := json.Unmarshal([]byte(questionsJSON), &batchReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid questions JSON format",
			"error":   err.Error(),
		})
	}

	// Validasi jumlah pertanyaan
	if len(batchReq.Questions) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "At least one question is required",
		})
	}

	if len(batchReq.Questions) > 100 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Maximum 100 questions per batch",
		})
	}

	// Ambil semua file yang diupload
	// File diidentifikasi dengan key: image_0, image_1, image_2, dst
	files := make(map[int]*multipart.FileHeader)
	for i := 0; i < 100; i++ {
		fileKey := fmt.Sprintf("image_%d", i)
		file, err := c.FormFile(fileKey)
		if err != nil {
			break // Tidak ada file lagi
		}
		files[i] = file
	}

	// Set user context untuk semua pertanyaan
	userID := c.Get("UserId").(string)
	createdBy := c.Get("CreatedBy").(string)

	// Proses setiap question
	var processedQuestions []dto.MultipleRequest

	for i, question := range batchReq.Questions {
		// Set basic info
		question.UserID = userID
		question.CreatedBy = createdBy

		// Generate ID unik
		question.ID = generateID() // Atau pakai util.GenerateRandomString()

		// Set default question type
		if question.QuestionType == "" {
			question.QuestionType = "multiple_choice"
		}

		// Cek apakah pertanyaan ini memiliki gambar
		// image_index = -1 atau tidak ada gambar
		// image_index = 0,1,2 dst merujuk ke file yang diupload
		if question.ImageIndex >= 0 && question.ImageIndex < len(files) {
			if file, exists := files[question.ImageIndex]; exists {
				// Validasi file
				if err := multiple.ValidateImageFile(file); err != nil {
					return c.JSON(http.StatusBadRequest, map[string]interface{}{
						"status":  "error",
						"message": fmt.Sprintf("Invalid image for question %d (index %d): %s", i, question.ImageIndex, err.Error()),
					})
				}

				// Upload file
				imageURL, err := multiple.UploadFile(file, question.ID)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]interface{}{
						"status":  "error",
						"message": fmt.Sprintf("Failed to upload image for question %d", i),
						"error":   err.Error(),
					})
				}
				question.ImageURL = imageURL
			}
		}
		// Jika ImageIndex = -1 atau tidak ada file, biarkan ImageURL kosong

		processedQuestions = append(processedQuestions, question)
	}

	// Panggil service batch
	batchResponse, err := h.serviceMultiple.CreateBatch(processedQuestions)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to create multiple questions",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   batchResponse,
	})
}

// Create single question (optional)
func (h *domainHandler) Create(c echo.Context) error {
	c.Request().ParseMultipartForm(10 << 20)

	var req dto.MultipleRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	userID := c.Get("UserId").(string)
	req.UserID = userID

	createdBy := c.Get("CreatedBy").(string)
	req.CreatedBy = createdBy

	// Handle file upload if exists
	var imageURL string
	if req.Image != nil {
		if err := multiple.ValidateImageFile(req.Image); err != nil {
			return c.JSON(400, map[string]interface{}{
				"status":  "error",
				"message": err.Error(),
			})
		}

		req.ID = generateID()
		url, err := multiple.UploadFile(req.Image, req.ID)
		if err != nil {
			return c.JSON(500, map[string]interface{}{
				"status":  "error",
				"message": "Failed to upload image",
				"error":   err.Error(),
			})
		}
		imageURL = url
	}

	req.ImageURL = imageURL

	result, err := h.serviceMultiple.Create(req)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Failed to create question",
			"error":   err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"data":   result,
	})
}

// GetByID get single question
func (h *domainHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(400, map[string]interface{}{
			"status":  "error",
			"message": "ID is required",
		})
	}

	question, err := h.serviceMultiple.GetByID(id)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"status":  "error",
			"message": "Question not found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"data":   question,
	})
}

// GetByDetailID get all questions by detail ID
func (h *domainHandler) GetByDetailID(c echo.Context) error {
	userID := c.Param("user_id")
	if userID == "" {
		return c.JSON(400, map[string]interface{}{
			"status":  "error",
			"message": "Detail ID is required",
		})
	}

	questions, err := h.serviceMultiple.GetByDetailID(userID)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Failed to get questions",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"data":   questions,
	})
}

// Update update question
func (h *domainHandler) Update(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(400, map[string]interface{}{
			"status":  "error",
			"message": "ID is required",
		})
	}

	c.Request().ParseMultipartForm(10 << 20)

	var req dto.MultipleUpdateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	req.ID = id

	// Handle file upload if new image provided
	if req.Image != nil {
		if err := multiple.ValidateImageFile(req.Image); err != nil {
			return c.JSON(400, map[string]interface{}{
				"status":  "error",
				"message": err.Error(),
			})
		}

		url, err := multiple.UploadFile(req.Image, id)
		if err != nil {
			return c.JSON(500, map[string]interface{}{
				"status":  "error",
				"message": "Failed to upload image",
			})
		}
		req.ImageURL = url
	}

	if err := h.serviceMultiple.Update(id, req); err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Failed to update question",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status":  "success",
		"message": "Question updated successfully",
	})
}

// Delete delete question
func (h *domainHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(400, map[string]interface{}{
			"status":  "error",
			"message": "ID is required",
		})
	}

	if err := h.serviceMultiple.Delete(id); err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Failed to delete question",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status":  "success",
		"message": "Question deleted successfully",
	})
}

// Helper function to generate ID (sesuaikan dengan util yang Anda punya)
func generateID() string {
	// Gunakan implementasi generate random string Anda
	// Contoh sederhana:
	return fmt.Sprintf("q_%d", time.Now().UnixNano())
}
