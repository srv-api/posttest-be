package multiple

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	dto "posttest-be/dto/assessment/multiple"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	util "github.com/srv-api/util/s"
)

const (
	maxUploadSize = 5 << 20
	uploadDir     = "uploads/multiple"
)

func (h *domainHandler) Create(c echo.Context) error {
	c.Request().ParseMultipartForm(maxUploadSize)

	var req dto.MultipleRequest

	// Bind form data
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	userID := c.Get("UserId").(string)
	req.UserID = userID

	detailId := c.Get("DetailId").(string)
	req.DetailID = detailId

	createdBy := c.Get("CreatedBy").(string)
	req.CreatedBy = createdBy

	// Handle file upload
	var imageURL string
	if req.Image != nil {
		// Validasi file
		if err := validateImageFile(req.Image); err != nil {
			return c.JSON(400, map[string]interface{}{
				"status":  "error",
				"message": err.Error(),
			})
		}

		// Upload file
		url, err := uploadFile(req.Image)
		if err != nil {
			return c.JSON(500, map[string]interface{}{
				"status":  "error",
				"message": "Failed to upload image",
				"error":   err.Error(),
			})
		}
		imageURL = url
	}

	// Set image URL ke request
	req.ImageURL = imageURL

	// Panggil service
	multiple, err := h.serviceMultiple.Create(req)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Failed to create multiple",
			"error":   err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"data":   multiple,
	})
}

// Validasi file gambar
func validateImageFile(file *multipart.FileHeader) error {
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

	// Cek ukuran file (max 5MB sudah di-set di form parsing)
	if file.Size > maxUploadSize {
		return fmt.Errorf("file too large. Max size: 5MB")
	}

	// Buka file untuk validasi konten
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open file")
	}
	defer src.Close()

	// Baca header file untuk deteksi tipe MIME (5 bytes cukup)
	header := make([]byte, 512)
	_, err = src.Read(header)
	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to read file header")
	}

	// Reset file pointer
	src.Seek(0, 0)

	// Deteksi MIME type
	mimeType := http.DetectContentType(header)
	if !strings.HasPrefix(mimeType, "image/") {
		return fmt.Errorf("file is not an image")
	}

	return nil
}

// Upload file ke server
func uploadFile(file *multipart.FileHeader) (string, error) {
	// Buat direktori jika belum ada
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Generate unique filename
	timestamp := time.Now().UnixNano()
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d_%s%s", timestamp, util.GenerateRandomString(), ext)
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

	// Return URL path (bisa disesuaikan dengan base URL aplikasi)
	imageURL := "/" + filePath // atau menggunakan base URL: "https://yourdomain.com/" + filePath

	return imageURL, nil
}
