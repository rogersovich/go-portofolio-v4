package upload

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/rogersovich/go-portofolio-v4/dto"
)

func GenerateAdditionalInfo(input dto.UploadFileInput, folder string) (fileName string, contentType string, fileSize int64) {
	rawFileName := input.FileHeader.Filename

	// Ekstrak ekstensi file (misal: .jpg, .png)
	ext := filepath.Ext(rawFileName)

	// Gunakan UUID
	uniqueID := uuid.New().String()

	// Nama file baru
	fileName = fmt.Sprintf("%s/%d_%s%s", folder, time.Now().Unix(), uniqueID, ext)
	contentType = input.FileHeader.Header.Get("Content-Type")
	fileSize = input.FileHeader.Size

	return fileName, contentType, fileSize
}
