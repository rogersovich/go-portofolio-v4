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

var folderNameProjectFeature = "project_feature"

func GetAllProjectFeatures(params dto.ProjectFeatureQueryParams) ([]dto.ProjectFeatureResponse, error) {
	db, _ := config.DB.DB()

	var (
		conditions []string
		args       []interface{}
	)

	query := `SELECT id, image_url, image_file_name, description, created_at FROM project_features`

	// 🔍 Filters

	filters := []utils.SQLFilter{
		{Column: "description", Value: params.Description, Op: "LIKE"},
	}

	if params.IsDelete == "N" || params.IsDelete == "" {
		filters = append(filters, utils.SQLFilter{Column: "deleted_at", Op: "IS NULL", Value: true})
	} else if params.IsDelete == "Y" {
		filters = append(filters, utils.SQLFilter{Column: "deleted_at", Op: "IS NOT NULL", Value: true})
	}

	conditions, args = utils.BuildSQLFilters(filters)

	// 📅 Date Range (created_from & created_to)
	utils.AddDateRangeFilter("created_at", params.CreatedFrom, params.CreatedTo, &conditions, &args)

	// Add WHERE clause
	query += utils.BuildWhereClause(conditions)

	// 🧭 Append order + pagination
	query += utils.BuildOrderAndPagination(params.Order, params.Sort, params.Page, params.Limit)

	// Quer Debug

	utils.Log.Debug("Query SQL:", query)
	utils.Log.Debug("Conditons SQL:", conditions)

	rows, err := db.Query(query, args...)
	if err != nil {
		utils.LogError(err.Error(), query)
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var project_features []models.ProjectFeature

	for rows.Next() {
		var rowProjectFeature models.ProjectFeature
		if err := rows.Scan(
			&rowProjectFeature.ID,
			&rowProjectFeature.ImageUrl,
			&rowProjectFeature.ImageFileName,
			&rowProjectFeature.Description,
			&rowProjectFeature.CreatedAt); err != nil {
			utils.LogWarning(err.Error(), query)
			return nil, err
		}
		project_features = append(project_features, rowProjectFeature)
	}

	var response []dto.ProjectFeatureResponse
	for _, rowProjectFeature := range project_features {
		response = append(response, dto.ProjectFeatureResponse{
			ID:            rowProjectFeature.ID,
			ImageURL:      rowProjectFeature.ImageUrl,
			ImageFileName: rowProjectFeature.ImageFileName,
			Description:   &rowProjectFeature.Description,
			CreatedAt:     rowProjectFeature.CreatedAt.Format("2006-01-02"),
		})
	}

	return response, nil
}

func GetProjectFeature(id int) (dto.ProjectFeatureSingleResponse, error) {
	var response models.ProjectFeature
	if err := config.DB.First(&response, id).Error; err != nil {
		return dto.ProjectFeatureSingleResponse{}, err
	}

	return dto.ProjectFeatureSingleResponse{
		ID:            response.ID,
		Description:   &response.Description,
		ImageURL:      response.ImageUrl,
		ImageFileName: response.ImageFileName,
		CreatedAt:     response.CreatedAt.Format("2006-01-02"),
	}, nil
}

func CreateProjectFeature(req dto.CreateProjectFeatureRequest, c *gin.Context) (
	result dto.ProjectFeatureSingleResponse,
	statusCode int,
	errFiels []utils.FieldError,
	err error) {
	// set field
	imageFieldName := "image_file"

	// Upload image_file
	imageData, imageErrs, imageUploadErr := uploadService.HandleUploadedFile(
		c,
		imageFieldName,
		folderNameProjectFeature,
		nil,                           // use default allowed extensions
		2*1024*1024,                   // max 2MB
		[]string{"extension", "size"}, // []string{"required", "extension", "size"}
	)

	if imageErrs != nil {
		err = fmt.Errorf("invalid %s", imageFieldName)
		return result, http.StatusBadRequest, imageErrs, err
	}

	if imageUploadErr != nil {
		err = fmt.Errorf("failed to upload %s", imageFieldName)
		return result, http.StatusInternalServerError, imageErrs, err
	}

	data := models.ProjectFeature{
		Description:   req.Description,
		ImageUrl:      imageData.FileURL,
		ImageFileName: &imageData.FileName,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		return result, http.StatusInternalServerError, nil, err
	}

	result = dto.ProjectFeatureSingleResponse{
		ID:            data.ID,
		Description:   &data.Description,
		ImageURL:      data.ImageUrl,
		ImageFileName: data.ImageFileName,
		CreatedAt:     data.CreatedAt.Format("2006-01-02"),
	}

	return result, http.StatusOK, nil, nil
}

func UpdateProjectFeature(req dto.UpdateProjectFeatureRequest, id int, c *gin.Context) (
	result dto.ProjectFeatureUpdateResponse,
	statusCode int,
	errFiels []utils.FieldError,
	err error) {
	// set field
	imageFieldName := "image_file"
	// 1. Fetch existing ProjectFeature data
	resProjectFeature, err := GetProjectFeature(id)
	if err != nil {
		return result, http.StatusNotFound, nil, err
	}

	// set oldPath
	oldPath := ""
	if resProjectFeature.ImageFileName != nil {
		oldPath = *resProjectFeature.ImageFileName
	}

	// 2. Get new file (if uploaded)
	_, err = c.FormFile(imageFieldName)
	var newFileURL string
	var newFileName string

	if err == nil {
		// Upload Process
		imageData, imageErrs, imageUploadErr := uploadService.HandleUploadedFile(
			c,
			imageFieldName,
			folderNameProjectFeature,
			nil,                           // use default allowed extensions
			2*1024*1024,                   // max 2MB
			[]string{"extension", "size"}, // []string{"required", "extension", "size"}
		)

		if imageErrs != nil {
			err = fmt.Errorf("invalid %s", imageFieldName)
			return result, http.StatusBadRequest, imageErrs, err
		}

		if imageUploadErr != nil {
			err = fmt.Errorf("failed to upload %s", imageFieldName)
			return result, http.StatusInternalServerError, imageErrs, err
		}

		newFileURL = imageData.FileURL
		newFileName = imageData.FileName
	} else {
		newFileURL = resProjectFeature.ImageURL // keep existing if not updated
		newFileName = *resProjectFeature.ImageFileName
	}

	data := models.ProjectFeature{
		Description:   req.Description,
		ImageUrl:      newFileURL,
		ImageFileName: &newFileName,
	}

	if err := config.DB.Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return result, http.StatusInternalServerError, nil, err
	}

	// 3. Optional: Delete old file from MinIO
	if oldPath != newFileName {
		err = uploadService.DeleteFromMinio(c.Request.Context(), oldPath) // ignore error or handle if needed
		if err != nil {
			utils.Log.Warn(err.Error())
		}
	}

	return dto.ProjectFeatureUpdateResponse{
		Description:   &req.Description,
		ImageURL:      data.ImageUrl,
		ImageFileName: data.ImageFileName,
	}, http.StatusOK, nil, nil
}

func DeleteProjectFeature(id int, c *gin.Context) (statusCode int, err error) {
	// 1. Fetch existing ProjectFeature data
	_, err = GetProjectFeature(id)
	if err != nil {
		return http.StatusNotFound, err
	}

	// 3. Delete data
	if err := config.DB.Delete(&models.ProjectFeature{}, id).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
