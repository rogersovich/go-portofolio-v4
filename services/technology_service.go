package services

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/rogersovich/go-portofolio-v4/config"
	"github.com/rogersovich/go-portofolio-v4/dto"
	"github.com/rogersovich/go-portofolio-v4/models"
	"github.com/rogersovich/go-portofolio-v4/utils"
)

func GetAllTechnologies(params dto.TechnologyQueryParams) ([]dto.TechnologyResponse, error) {
	db, _ := config.DB.DB()

	var (
		conditions []string
		args       []interface{}
	)

	query := `SELECT id, name, logo_url, description_html, is_major, created_at FROM technologies`

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
		if err := rows.Scan(&tech.ID, &tech.Name, &tech.LogoURL, &tech.DescriptionHTML, &tech.IsMajor, &tech.CreatedAt); err != nil {
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
			Logo:            tech.LogoURL,
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
		CreatedAt:       tech.CreatedAt.Format("2006-01-02"),
	}, nil
}

func CreateTechnology(req dto.CreateTechnologyRequest) (dto.TechnologySingleResponse, error) {
	data := models.Technology{
		Name:            req.Name,
		DescriptionHTML: req.DescriptionHTML,
		LogoURL:         req.LogoURL,
		IsMajor:         strings.ToUpper(req.IsMajor) == "Y",
	}

	var result dto.TechnologySingleResponse

	if err := config.DB.Create(&data).Error; err != nil {
		return result, err
	}

	result = dto.TechnologySingleResponse{
		ID:              data.ID,
		Name:            data.Name,
		DescriptionHTML: data.DescriptionHTML,
		LogoURL:         data.LogoURL,
		IsMajor:         utils.BoolToYN(data.IsMajor),
		CreatedAt:       data.CreatedAt.Format("2006-01-02"),
	}

	return result, nil
}

func UpdateTechnology(req dto.UpdateTechnologyRequest, id int) (dto.TechnologyUpdateSingleResponse, error) {
	data := models.Technology{
		Name:            req.Name,
		DescriptionHTML: req.DescriptionHTML,
		LogoURL:         req.LogoURL,
		IsMajor:         strings.ToUpper(req.IsMajor) == "Y",
	}

	var result dto.TechnologyUpdateSingleResponse

	if err := config.DB.Where("id = ?", id).Updates(&data).Error; err != nil {
		return result, err
	}

	result = dto.TechnologyUpdateSingleResponse{
		Name:            data.Name,
		DescriptionHTML: data.DescriptionHTML,
		LogoURL:         data.LogoURL,
		IsMajor:         utils.BoolToYN(data.IsMajor),
	}

	return result, nil
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
