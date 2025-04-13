package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/config"
	"github.com/rogersovich/go-portofolio-v4/dto"
	"github.com/rogersovich/go-portofolio-v4/models"
)

func GetProjectTechnology(id int) (dto.ProjectTechnologySingleResponse, error) {
	var response models.ProjectTechnology
	if err := config.DB.First(&response, id).Error; err != nil {
		return dto.ProjectTechnologySingleResponse{}, err
	}

	return dto.ProjectTechnologySingleResponse{
		ID:           response.ID,
		ProjectID:    response.ProjectID,
		TechnologyID: response.TechnologyID,
		CreatedAt:    response.CreatedAt.Format("2006-01-02"),
	}, nil
}

func CreateProjectTechnology(req dto.CreateProjectTechnologyRequest) (result dto.ProjectTechnologySingleResponse, err error) {
	data := models.ProjectTechnology{
		ProjectID:    req.ProjectID,
		TechnologyID: req.TechnologyID,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		return result, err
	}

	result = dto.ProjectTechnologySingleResponse{
		ID:           data.ID,
		ProjectID:    data.ProjectID,
		TechnologyID: data.TechnologyID,
		CreatedAt:    data.CreatedAt.Format("2006-01-02"),
	}

	return result, nil
}

func UpdateProjectTechnology(req dto.UpdateProjectTechnologyRequest, id int) (result dto.ProjectTechnologyUpdateResponse, err error) {
	data := models.ProjectTechnology{
		ProjectID:    req.ProjectID,
		TechnologyID: req.TechnologyID,
	}

	if err := config.DB.Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return result, err
	}

	result = dto.ProjectTechnologyUpdateResponse{
		ProjectID:    req.ProjectID,
		TechnologyID: req.TechnologyID,
	}

	return result, nil
}

func DeleteProjectTechnology(id int, c *gin.Context) (statusCode int, err error) {
	// 1. Fetch existing data
	_, err = GetProjectTechnology(id)
	if err != nil {
		return http.StatusNotFound, err
	}

	// 3. Delete data
	if err := config.DB.Delete(&models.ProjectTechnology{}, id).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
