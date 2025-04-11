package upload

import (
	"context"
	"fmt"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/rogersovich/go-portofolio-v4/dto"
	"github.com/rogersovich/go-portofolio-v4/utils"
)

func UploadFile(ctx context.Context, input dto.UploadFileInput, folder string) (*dto.UploadResponse, error) {
	endpoint := utils.GetEnv("MINIO_ENDPOINT")
	bucketName := utils.GetEnv("MINIO_BUCKET")

	// Init client
	minioClient, err := GenerateMinioClient()
	if err != nil {
		return nil, err
	}

	// Generate additional info
	fileName, contentType, fileSize := GenerateAdditionalInfo(input, folder)

	// Upload to MinIO
	_, err = minioClient.PutObject(ctx, bucketName, fileName, input.File, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return nil, err
	}

	// Generate public URL (jika public)
	fileURL := BuildMinioURL(endpoint, bucketName, fileName)

	// Generate response
	avatar := dto.UploadResponse{
		FileURL:  fileURL,
		FileName: fileName,
	}

	return &avatar, nil
}

func HandleUploadedFile(
	ctx *gin.Context,
	fieldName string,
	folderName string,
	allowedExt []string,
	maxSize int64,
	validationCheck []string,
) (dto.UploadResponse, []utils.FieldError, error) {

	if len(validationCheck) == 0 {
		validationCheck = []string{"required", "extension", "size"}
	}

	// Step 1: Get the file
	file, err := ctx.FormFile(fieldName)
	if err != nil && slices.Contains(validationCheck, "required") {
		return dto.UploadResponse{}, utils.GenerateFieldErrorResponse(fieldName, fmt.Sprintf("%s is required", fieldName)), nil
	}

	// Step 2: Validate extension
	errExt := ValidateExtension(file.Filename, allowedExt)
	if errExt != nil && slices.Contains(validationCheck, "extension") {
		return dto.UploadResponse{}, errExt, nil
	}

	// Step 3: Validate size
	if file.Size > maxSize && slices.Contains(validationCheck, "size") {
		return dto.UploadResponse{}, utils.GenerateFieldErrorResponse(fieldName, fmt.Sprintf("%s exceeds max size", fieldName)), nil
	}

	// Step 4: Open file
	openedFile, err := file.Open()
	if err != nil {
		return dto.UploadResponse{}, nil, err
	}
	defer openedFile.Close()

	// Step 5: Upload to MinIO
	payload := dto.UploadFileInput{
		FileHeader: file,
		File:       openedFile,
	}
	uploadedData, err := UploadFile(ctx.Request.Context(), payload, folderName)
	if err != nil {
		return dto.UploadResponse{}, nil, err
	}

	return *uploadedData, nil, nil
}

func DeleteFromMinio(ctx context.Context, objectPath string) error {
	bucketName := utils.GetEnv("MINIO_BUCKET")

	// Init client
	minioClient, err := GenerateMinioClient()
	if err != nil {
		return err
	}

	return minioClient.RemoveObject(ctx, bucketName, objectPath, minio.RemoveObjectOptions{})
}
