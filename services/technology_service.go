package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/config"
	"github.com/rogersovich/go-portofolio-v4/dto"
	"github.com/rogersovich/go-portofolio-v4/models"
	uploadService "github.com/rogersovich/go-portofolio-v4/services/upload"
	"github.com/rogersovich/go-portofolio-v4/utils"
)

var folderNameTechnology = "technology"

func GetAllTechnologies(params dto.TechnologyQueryParams) ([]dto.TechnologyResponse, error) {
	db, _ := config.DB.DB()

	var (
		conditions []string
		args       []interface{}
	)

	query := `SELECT id, name, logo_url, logo_file_name, description_html, is_major, created_at FROM technologies`

	// 🔍 Filters

	filters := []utils.SQLFilter{
		{Column: "name", Value: params.FilterName, Op: "LIKE"},
		{Column: "description_html", Value: params.FilterDesc, Op: "LIKE"},
	}

	// Handle boolean flags
	if params.IsMajor == "Y" {
		filters = append(filters, utils.SQLFilter{Column: "is_major", Op: "=", Value: true})
	} else if params.IsMajor == "N" {
		filters = append(filters, utils.SQLFilter{Column: "is_major", Op: "=", Value: false})
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

	var technologies []models.Technology

	for rows.Next() {
		var tech models.Technology
		if err := rows.Scan(&tech.ID, &tech.Name, &tech.LogoURL, &tech.LogoFileName, &tech.DescriptionHTML, &tech.IsMajor, &tech.CreatedAt); err != nil {
			utils.LogWarning(err.Error(), query)
			return nil, err
		}
		technologies = append(technologies, tech)
	}

	var response []dto.TechnologyResponse
	for _, tech := range technologies {
		response = append(response, dto.TechnologyResponse{
			ID:              tech.ID,
			Name:            tech.Name,
			LogoURL:         tech.LogoURL,
			LogoFileName:    tech.LogoFileName,
			DescriptionHTML: tech.DescriptionHTML,
			Major:           utils.BoolToYN(tech.IsMajor),
			CreatedAt:       tech.CreatedAt.Format("2006-01-02"),
		})
	}

	return response, nil
}

func GetTechnology(id int) (dto.TechnologySingleResponse, error) {
	var tech models.Technology
	if err := config.DB.First(&tech, id).Error; err != nil {
		return dto.TechnologySingleResponse{}, err
	}

	return dto.TechnologySingleResponse{
		ID:              tech.ID,
		Name:            tech.Name,
		DescriptionHTML: tech.DescriptionHTML,
		LogoURL:         tech.LogoURL,
		LogoFileName:    tech.LogoFileName,
		IsMajor:         utils.BoolToYN(tech.IsMajor),
		CreatedAt:       tech.CreatedAt.Format("2006-01-02"),
	}, nil
}

func CreateTechnology(req dto.CreateTechnologyRequest, c *gin.Context) (result dto.TechnologySingleResponse, statusCode int, errFiels []utils.FieldError, err error) {
	// Upload logo_file
	logoFile, logoErrs, logoUploadErr := uploadService.HandleUploadedFile(
		c,
		"logo_file",
		folderNameTechnology,
		nil,         // use default allowed extensions
		2*1024*1024, // max 2MB
		nil,         // []string{"required", "extension", "size"}
	)

	if logoErrs != nil {
		err = fmt.Errorf("invalid avatar_file")
		return result, http.StatusBadRequest, logoErrs, err
	}

	if logoUploadErr != nil {
		err = fmt.Errorf("failed to upload avatar_file")
		return result, http.StatusInternalServerError, logoErrs, err
	}

	data := models.Technology{
		Name:            req.Name,
		DescriptionHTML: &req.DescriptionHTML,
		LogoURL:         &logoFile.FileURL,
		LogoFileName:    &logoFile.FileName,
		IsMajor:         strings.ToUpper(req.IsMajor) == "Y",
	}

	if err := config.DB.Create(&data).Error; err != nil {
		return result, http.StatusInternalServerError, nil, err
	}

	result = dto.TechnologySingleResponse{
		ID:              data.ID,
		Name:            data.Name,
		DescriptionHTML: data.DescriptionHTML,
		LogoURL:         data.LogoURL,
		LogoFileName:    data.LogoFileName,
		IsMajor:         utils.BoolToYN(data.IsMajor),
		CreatedAt:       data.CreatedAt.Format("2006-01-02"),
	}

	return result, http.StatusOK, nil, nil
}

func UpdateTechnology(req dto.UpdateTechnologyRequest, id int, c *gin.Context) (result dto.TechnologyUpdateSingleResponse, statusCode int, errFiels []utils.FieldError, err error) {
	// 1. Fetch existing about data
	oldData, err := GetTechnology(id)
	if err != nil {
		return result, http.StatusNotFound, nil, err
	}

	// set oldPath
	oldPath := oldData.LogoFileName

	// 2. Get new file (if uploaded)
	_, err = c.FormFile("logo_file")
	var newFileURL string
	var newFileName string

	if err == nil {
		// Upload logo_file
		logoData, logo_fileErrs, logoUploadErr := uploadService.HandleUploadedFile(
			c,
			"logo_file",
			folderNameTechnology,
			nil,                           // use default allowed extensions
			2*1024*1024,                   // max 2MB
			[]string{"extension", "size"}, // []string{"required", "extension", "size"}
		)

		if logo_fileErrs != nil {
			err = fmt.Errorf("invalid logo_file")
			return result, http.StatusBadRequest, logo_fileErrs, err
		}

		if logoUploadErr != nil {
			err = fmt.Errorf("failed to upload logo_file")
			return result, http.StatusInternalServerError, logo_fileErrs, err
		}

		newFileURL = logoData.FileURL
		newFileName = logoData.FileName
	} else {
		newFileURL = *oldData.LogoURL // keep existing if not updated
		newFileName = *oldData.LogoFileName
	}

	data := models.Technology{
		Name:            req.Name,
		DescriptionHTML: &req.DescriptionHTML,
		LogoURL:         &newFileURL,
		LogoFileName:    &newFileName,
		IsMajor:         strings.ToUpper(req.IsMajor) == "Y",
	}

	if err := config.DB.Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return result, http.StatusInternalServerError, nil, err
	}

	// 3. Optional: Delete old file from MinIO
	if oldPath != &newFileName {
		err = uploadService.DeleteFromMinio(c.Request.Context(), *oldPath) // ignore error or handle if needed
		if err != nil {
			utils.Log.Warn(err.Error())
		}
	}

	return dto.TechnologyUpdateSingleResponse{
		Name:            data.Name,
		DescriptionHTML: data.DescriptionHTML,
		LogoURL:         &newFileURL,
		LogoFileName:    &newFileName,
		IsMajor:         utils.BoolToYN(data.IsMajor),
	}, http.StatusOK, nil, nil
}

func DeleteTechnology(id int) (dto.TechnologyDeleteSingleResponse, error) {
	db, _ := config.DB.DB()

	table := "technologies"

	// Check if the row exists and not already soft-deleted
	selectQuery := `SELECT id, name, deleted_at FROM ` + table + ` WHERE id = ? AND deleted_at IS NULL`

	var result dto.TechnologyDeleteSingleResponse

	// Query single row
	err := db.QueryRow(selectQuery, id).Scan(&result.ID, &result.Name, &result.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.LogError(err.Error(), "")
			return result, fmt.Errorf("technology with id %d not found or already deleted", id)
		}
		utils.LogError(err.Error(), selectQuery)
		return result, fmt.Errorf("query error: %w", err)
	}

	// Perform soft delete
	updateQuery := `UPDATE ` + table + ` SET deleted_at = CURRENT_TIMESTAMP WHERE id = ?`

	_, err = db.Exec(updateQuery, id)
	if err != nil {
		utils.LogError(err.Error(), updateQuery)
		return result, fmt.Errorf("failed to delete technology: %w", err)
	}

	return result, nil
}
