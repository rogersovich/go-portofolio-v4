package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/config"
	"github.com/rogersovich/go-portofolio-v4/dto"
	"github.com/rogersovich/go-portofolio-v4/models"
	"github.com/rogersovich/go-portofolio-v4/utils"
)

func GetAllStatistics(params dto.StatisticQueryParams) ([]dto.StatisticResponse, error) {
	db, _ := config.DB.DB()

	var (
		conditions []string
		args       []interface{}
	)

	query := `SELECT id, likes, views, type, created_at FROM statistics`

	// 🔍 Filters

	filters := []utils.SQLFilter{
		{Column: "type", Value: params.Type, Op: "="},
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

	var statistics []models.Statistic

	for rows.Next() {
		var tech models.Statistic
		if err := rows.Scan(&tech.ID, &tech.Likes, &tech.Views, &tech.Type, &tech.CreatedAt); err != nil {
			utils.LogWarning(err.Error(), query)
			return nil, err
		}
		statistics = append(statistics, tech)
	}

	var response []dto.StatisticResponse
	for _, tech := range statistics {
		response = append(response, dto.StatisticResponse{
			ID:        tech.ID,
			Likes:     tech.Likes,
			Views:     tech.Views,
			Type:      tech.Type,
			CreatedAt: tech.CreatedAt.Format("2006-01-02"),
		})
	}

	return response, nil
}

func GetStatistic(id int) (dto.StatisticSingleResponse, error) {
	var response models.Statistic
	if err := config.DB.First(&response, id).Error; err != nil {
		return dto.StatisticSingleResponse{}, err
	}

	return dto.StatisticSingleResponse{
		ID:        response.ID,
		Likes:     *response.Likes,
		Views:     *response.Views,
		Type:      response.Type,
		CreatedAt: response.CreatedAt.Format("2006-01-02"),
	}, nil
}

func CreateStatistic(req dto.CreateStatisticRequest) (result dto.StatisticSingleResponse, err error) {
	data := models.Statistic{
		Likes: req.Likes,
		Views: req.Views,
		Type:  req.Type,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		return result, err
	}

	result = dto.StatisticSingleResponse{
		ID:        data.ID,
		Likes:     *data.Likes,
		Views:     *data.Views,
		Type:      data.Type,
		CreatedAt: data.CreatedAt.Format("2006-01-02"),
	}

	return result, nil
}

func UpdateStatistic(req dto.UpdateStatisticRequest, id int) (result dto.StatisticUpdateResponse, err error) {
	data := models.Statistic{
		Likes: req.Likes,
		Views: req.Views,
		Type:  req.Type,
	}

	if err := config.DB.Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return result, err
	}

	result = dto.StatisticUpdateResponse{
		Likes: *data.Likes,
		Views: *data.Views,
		Type:  data.Type,
	}

	return result, nil
}

func DeleteStatistic(id int, c *gin.Context) (statusCode int, err error) {
	// 1. Fetch existing data
	_, err = GetStatistic(id)
	if err != nil {
		return http.StatusNotFound, err
	}

	// 3. Delete data
	if err := config.DB.Delete(&models.Statistic{}, id).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
