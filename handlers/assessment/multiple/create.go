package multiple

import (
	"net/http"
	dto "posttest-be/dto/assessment/multiple"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) Create(c echo.Context) error {
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

	detailID := c.Get("DetailId").(string)
	req.DetailID = detailID

	createdBy := c.Get("CreatedBy").(string)
	req.CreatedBy = createdBy

	multiple, err := h.serviceMultiple.Create(req)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Failed to create question",
			"error":   err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"data":   multiple,
	})
}

// CreateBatch creates multiple questions with base64 images
func (h *domainHandler) CreateBatch(c echo.Context) error {
	var batchReq dto.MultipleBatchRequest

	// Bind JSON body
	if err := c.Bind(&batchReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validasi minimal 1 pertanyaan
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

	// Set user context untuk semua pertanyaan
	userID := c.Get("UserId").(string)
	detailID := c.Get("DetailId").(string)
	createdBy := c.Get("CreatedBy").(string)

	for i := range batchReq.Questions {
		batchReq.Questions[i].UserID = userID
		batchReq.Questions[i].DetailID = detailID
		batchReq.Questions[i].CreatedBy = createdBy
	}

	// Panggil service batch
	batchResponse, err := h.serviceMultiple.CreateBatch(batchReq.Questions)
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

// GetByID get single question by ID
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
			"error":   err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"data":   question,
	})
}

// GetByDetailID get all questions by detail ID
func (h *domainHandler) GetByDetailID(c echo.Context) error {
	detailID := c.Param("detail_id")
	if detailID == "" {
		return c.JSON(400, map[string]interface{}{
			"status":  "error",
			"message": "Detail ID is required",
		})
	}

	questions, err := h.serviceMultiple.GetByDetailID(detailID)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Failed to get questions",
			"error":   err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"data":   questions,
	})
}

// Update update existing question
func (h *domainHandler) Update(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(400, map[string]interface{}{
			"status":  "error",
			"message": "ID is required",
		})
	}

	var req dto.MultipleUpdateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	req.ID = id

	if err := h.serviceMultiple.Update(id, req); err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Failed to update question",
			"error":   err.Error(),
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
			"error":   err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status":  "success",
		"message": "Question deleted successfully",
	})
}
