package services

import (
	"fmt"
	"strings"

	"github.com/rogersovich/go-portofolio-v4/config"
	"github.com/rogersovich/go-portofolio-v4/dto"
	"github.com/rogersovich/go-portofolio-v4/models"
	"github.com/rogersovich/go-portofolio-v4/utils"
)

type TechnologyQueryParams struct {
	Sort        string
	Order       string
	FilterName  string
	FilterDesc  string
	IsMajor     string // "Y" or "N"
	IsDelete    string // "Y" or "N"
	CreatedFrom string
	CreatedTo   string
	Page        int
	Limit       int
}

func GetAllTechnologies(params TechnologyQueryParams) ([]dto.TechnologyResponse, error) {
	fmt.Println("params", params)
	db, _ := config.DB.DB()

	var (
		conditions []string
		args       []interface{}
	)

	query := `SELECT id, name, logo_url, description_html, is_major, created_at FROM technologies`

	// 🔍 Filters
	if params.FilterName != "" {
		conditions = append(conditions, "name LIKE ?")
		args = append(args, "%"+params.FilterName+"%")
	}
	if params.FilterDesc != "" {
		conditions = append(conditions, "description_html LIKE ?")
		args = append(args, "%"+params.FilterDesc+"%")
	}
	// ✅ IsMajor filter
	if params.IsMajor == "Y" {
		conditions = append(conditions, "is_major = TRUE")
	} else if params.IsMajor == "N" {
		conditions = append(conditions, "is_major = FALSE")
	}
	// ✅ IsDelete filter
	if params.IsDelete == "N" {
		conditions = append(conditions, "deleted_at IS NULL")
	} else if params.IsDelete == "Y" {
		conditions = append(conditions, "deleted_at IS NOT NULL")
	}

	// 📅 Date Range (created_from & created_to)
	utils.AddDateRangeFilter("created_at", params.CreatedFrom, params.CreatedTo, &conditions, &args)

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	fmt.Println(conditions, args)

	// 🧭 Sorting
	order := "id"
	if params.Order != "" {
		order = params.Order
	}
	sort := "ASC"
	if strings.ToUpper(params.Sort) == "DESC" {
		sort = "DESC"
	}
	query += fmt.Sprintf(" ORDER BY %s %s", order, sort)

	// 🧮 Pagination
	limit := params.Limit
	page := params.Page
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	// Quer Debug

	utils.Log.Debug("Built Query:", query)

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
			Major:           tech.IsMajor,
			CreatedAt:       tech.CreatedAt.Format("2006-01-02"),
		})
	}

	return response, nil
}

func GetTechnology(id int) (dto.TechnologySingleResponse, error) {
	var result dto.TechnologySingleResponse
	var tech models.Technology

	if err := config.DB.Table("technologies").
		Where("id = ?", id).
		Select("id", "name", "description_html", "logo_url", "is_major", "created_at").
		First(&tech).Error; err != nil {
		return result, err
	}

	fmt.Println(tech.IsMajor)

	result = dto.TechnologySingleResponse{
		ID:              tech.ID,
		Name:            tech.Name,
		DescriptionHTML: tech.DescriptionHTML,
		LogoURL:         tech.LogoURL,
		IsMajor:         utils.BoolToYN(tech.IsMajor),
		CreatedAt:       tech.CreatedAt.Format("2006-01-02"),
	}

	return result, nil

}
