package services

import (
	"fmt"

	"github.com/rogersovich/go-portofolio-v4/config"
	"github.com/rogersovich/go-portofolio-v4/dto"
	"github.com/rogersovich/go-portofolio-v4/models"
	"github.com/rogersovich/go-portofolio-v4/utils"
)

func GetAllAbouts(params dto.AboutQueryParams) ([]dto.AboutResponse, error) {
	db, _ := config.DB.DB()

	var (
		conditions []string
		args       []interface{}
	)

	query := `SELECT id, title, avatar_url, description_html, created_at FROM abouts`

	// 🔍 Filters

	filters := []utils.SQLFilter{
		{Column: "title", Value: params.Title, Op: "LIKE"},
		{Column: "description_html", Value: params.Description, Op: "LIKE"},
	}

	if params.IsDelete == "N" {
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

	var abouts []models.About

	for rows.Next() {
		var rowAbout models.About
		if err := rows.Scan(&rowAbout.ID, &rowAbout.Title, &rowAbout.AvatarUrl, &rowAbout.DescriptionHTML, &rowAbout.CreatedAt); err != nil {
			utils.LogWarning(err.Error(), query)
			return nil, err
		}
		abouts = append(abouts, rowAbout)
	}

	var response []dto.AboutResponse
	for _, rowAbout := range abouts {
		response = append(response, dto.AboutResponse{
			ID:              rowAbout.ID,
			Title:           rowAbout.Title,
			AvatarURL:       rowAbout.AvatarUrl,
			DescriptionHTML: rowAbout.DescriptionHTML,
			CreatedAt:       rowAbout.CreatedAt.Format("2006-01-02"),
		})
	}

	return response, nil
}

func CreateAbout(req dto.CreateAboutRequest) (dto.AboutSingleResponse, error) {
	tech := models.About{
		Title:           req.Title,
		DescriptionHTML: req.DescriptionHTML,
	}

	var result dto.AboutSingleResponse

	if err := config.DB.Create(&tech).Error; err != nil {
		return result, err
	}

	result = dto.AboutSingleResponse{
		ID:              tech.ID,
		Title:           tech.Title,
		DescriptionHTML: tech.DescriptionHTML,
		CreatedAt:       tech.CreatedAt.Format("2006-01-02"),
	}

	return result, nil
}
