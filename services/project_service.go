package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/config"
	"github.com/rogersovich/go-portofolio-v4/dto"
	"github.com/rogersovich/go-portofolio-v4/models"
	uploadService "github.com/rogersovich/go-portofolio-v4/services/upload"
	"github.com/rogersovich/go-portofolio-v4/utils"
)

var folderNameProject = "project"

func CheckProjectTechnology(technology_ids []string) error {
	db, _ := config.DB.DB()

	// Build placeholder string like "?,?,?,?"
	inClause, args := utils.BuildSQLInClause(technology_ids)

	// Build query
	query := fmt.Sprintf("SELECT id FROM technologies WHERE id IN (%s)", inClause)

	// Execute query
	rows, err := db.Query(query, args...)
	if err != nil {
		utils.LogError(err.Error(), query)
		return fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	// Check if all technology_ids exist
	count := 0
	for rows.Next() {
		count++
	}

	if count != len(technology_ids) {
		return fmt.Errorf("invalid technology_ids")
	}

	return nil
}

func CreateProject(req dto.CreateProjectRequest, c *gin.Context) (result dto.ProjectResponse, statusCode int, errFiels []utils.FieldError, err error) {
	// set field
	imageFieldName := "image_file"
	// Upload avatar_file
	imageData, imageErrs, imageUploadErr := uploadService.HandleUploadedFile(
		c,
		imageFieldName,
		folderNameProject,
		nil,         // use default allowed extensions
		2*1024*1024, // max 2MB
		nil,         // []string{"required", "extension", "size"}
	)

	if imageErrs != nil {
		err = fmt.Errorf("invalid %s", imageFieldName)
		return result, http.StatusBadRequest, imageErrs, err
	}

	if imageUploadErr != nil {
		err = fmt.Errorf("failed to upload %s", imageFieldName)
		return result, http.StatusInternalServerError, imageErrs, err
	}

	var publishedAt *time.Time

	if req.IsPublihed == "Y" {
		now := time.Now()
		publishedAt = &now
	} else {
		publishedAt = nil
	}

	data := models.Project{
		Title:         req.Title,
		Description:   req.Description,
		ImageURL:      &imageData.FileURL,
		ImageFileName: &imageData.FileName,
		RepositoryURL: req.RepositoryURL,
		Summary:       req.Summary,
		Status:        req.Status,
		PublishedAt:   publishedAt,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		return result, http.StatusInternalServerError, nil, err
	}

	result = dto.ProjectResponse{
		ID:            data.ID,
		Title:         data.Title,
		ImageURL:      *data.ImageURL,
		ImageFileName: *data.ImageFileName,
		CreatedAt:     data.CreatedAt.Format("2006-01-02"),
	}

	return result, http.StatusOK, nil, nil
}
