package services

import (
	"fmt"
	"net/http"
	"strconv"
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

func CheckProjectContentImage(content_images []string) error {
	db, _ := config.DB.DB()

	// Build placeholder string like "?,?,?,?"
	inClause, args := utils.BuildSQLInClause(content_images)

	// Build query
	query := fmt.Sprintf("SELECT id FROM project_content_images WHERE image_url IN (%s)", inClause)

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

	if count != len(content_images) {
		return fmt.Errorf("invalid content_images")
	}

	return nil
}

func CreateProject(req dto.CreateProjectRequest, c *gin.Context) (result dto.ProjectCreateResponse, statusCode int, errFiels []utils.FieldError, err error) {
	tx := config.DB.Begin()

	var likes int = 0
	var views int = 0
	resStatistic := models.Statistic{
		Likes: likes,
		Views: views,
		Type:  "Project",
	}

	if err := config.DB.Create(&resStatistic).Error; err != nil {
		tx.Rollback()
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return result, http.StatusInternalServerError, nil, err
	}

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
		StatisticId:   int(resStatistic.ID),
	}

	if err := config.DB.Create(&data).Error; err != nil {
		tx.Rollback()
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return result, http.StatusInternalServerError, nil, err
	}

	// INSERT to TABLE PROJECT TECHONOLOGIES
	var technologies []models.ProjectTechnology

	for _, technology_id := range req.TechnologyIds {
		TechnologyId, _ := strconv.Atoi(technology_id)
		technologies = append(technologies, models.ProjectTechnology{
			ProjectID:    1,
			TechnologyID: TechnologyId,
		})
	}

	if err := config.DB.Table("project_technologies").Create(&technologies).Error; err != nil {
		tx.Rollback()
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return result, http.StatusInternalServerError, nil, err
	}

	// UPDATE to TABLE PROJECT CONTENT IMAGES
	err = config.DB.Model(&models.ProjectContentImage{}).
		Where("image_url IN ?", req.ContentImages).
		Update("project_id", data.ID).Error

	if err != nil {
		tx.Rollback()
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return result, http.StatusInternalServerError, nil, err
	}

	tx.Commit()
	result = dto.ProjectCreateResponse{
		ID:            data.ID,
		Title:         data.Title,
		Description:   data.Description,
		ImageURL:      *data.ImageURL,
		ImageFileName: *data.ImageFileName,
		RepositoryURL: data.RepositoryURL,
		Summary:       data.Summary,
		Status:        data.Status,
		StatisticId:   int(resStatistic.ID),
		CreatedAt:     data.CreatedAt.Format("2006-01-02"),
	}

	return result, http.StatusOK, nil, nil
}
