package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/config"
	"github.com/rogersovich/go-portofolio-v4/dto"
	"github.com/rogersovich/go-portofolio-v4/models"
	uploadService "github.com/rogersovich/go-portofolio-v4/services/upload"
	"github.com/rogersovich/go-portofolio-v4/utils"
)

var folderNameProjectContentImage = "project"

func GetProjectContentImage(id int) (dto.ProjectContentImageResponse, error) {
	var response models.ProjectContentImage

	err := config.DB.
		Joins("LEFT JOIN projects ON projects.id = project_content_images.project_id").
		First(&response, id).Error

	if err != nil {
		return dto.ProjectContentImageResponse{}, err
	}

	return dto.ProjectContentImageResponse{
		ID:            response.ID,
		ProjectId:     response.ProjectId,
		ProjectName:   &response.Project.Title,
		IsUsed:        utils.BoolToYN(response.IsUsed),
		ImageURL:      response.ImageUrl,
		ImageFileName: response.ImageFileName,
		CreatedAt:     response.CreatedAt.Format("2006-01-02"),
	}, nil
}

func CreateProjectContentImage(req dto.CreateProjectContentImageRequest, c *gin.Context) (result dto.ProjectContentImageResponse, statusCode int, errFiels []utils.FieldError, err error) {

	// Upload image_file
	imageData, imageErrs, imageUploadErr := uploadService.HandleUploadedFile(
		c,
		"image_file",
		folderNameProjectContentImage,
		nil,         // use default allowed extensions
		2*1024*1024, // max 2MB
		nil,         // []string{"required", "extension", "size"}
	)

	if imageErrs != nil {
		err = fmt.Errorf("invalid image_file")
		return result, http.StatusBadRequest, imageErrs, err
	}

	if imageUploadErr != nil {
		err = fmt.Errorf("failed to upload image_file")
		return result, http.StatusInternalServerError, imageErrs, err
	}

	projectId := utils.ParseNullableInt(req.ProjectId)

	data := models.ProjectContentImage{
		ProjectId:     projectId,
		IsUsed:        req.IsUsed == "Y",
		ImageUrl:      imageData.FileURL,
		ImageFileName: imageData.FileName,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		return result, http.StatusInternalServerError, nil, err
	}

	result = dto.ProjectContentImageResponse{
		ID:            data.ID,
		ProjectId:     data.ProjectId,
		IsUsed:        utils.BoolToYN(data.IsUsed),
		ImageURL:      data.ImageUrl,
		ImageFileName: data.ImageFileName,
		CreatedAt:     data.CreatedAt.Format("2006-01-02"),
	}

	return result, http.StatusOK, nil, nil
}
