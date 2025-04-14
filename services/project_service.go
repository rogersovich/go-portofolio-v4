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

func GetAllProjects(params dto.ProjectQueryParams) ([]dto.ProjectGetAllDTO, error) {
	db, _ := config.DB.DB()

	var (
		conditions []string
		args       []interface{}
	)

	query := `
		SELECT 
			p.id,
			p.title,
			p.status,
			p.summary,
			p.image_url,
			p.repository_url,
			p.published_at,
			t.id AS tech_id,
			t.name AS tech_name
		FROM projects p
		JOIN project_technologies pt ON pt.project_id = p.id
		JOIN technologies t ON t.id = pt.technology_id
	`

	// 🔍 Filters

	filters := []utils.SQLFilter{
		{Column: "p.title", Value: params.Title, Op: "LIKE"},
		{Column: "p.description", Value: params.Description, Op: "LIKE"},
	}

	if params.IsDelete == "N" || params.IsDelete == "" {
		filters = append(filters, utils.SQLFilter{Column: "p.deleted_at", Op: "IS NULL", Value: true})
	} else if params.IsDelete == "Y" {
		filters = append(filters, utils.SQLFilter{Column: "p.deleted_at", Op: "IS NOT NULL", Value: true})
	}

	conditions, args = utils.BuildSQLFilters(filters)

	// 📅 Date Range (created_from & created_to)
	utils.AddDateRangeFilter("created_at", params.CreatedFrom, params.CreatedTo, &conditions, &args)

	// Add WHERE clause
	query += utils.BuildWhereClause(conditions)

	// 🧭 Append order + pagination
	query += utils.BuildOrderAndPagination(params.Order, params.Sort, params.Page, params.Limit)

	// Query Debug
	utils.Log.Debug("Query SQL:", query)
	utils.Log.Debug("Conditons SQL:", conditions)

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	// Mapping result
	projectMap := map[int]*dto.ProjectGetAllDTO{}

	for rows.Next() {
		var row dto.ProjectRawResponse
		if err := rows.Scan(
			&row.ID,
			&row.Title,
			&row.Status,
			&row.Summary,
			&row.ImageURL,
			&row.RepositoryURL,
			&row.PublishedAt,
			&row.TechID,
			&row.TechName); err != nil {
			utils.LogWarning(err.Error(), query)
			return nil, err
		}

		projectID := int(row.ID)

		//? "Comma-ok" itu fitur spesial
		_, exists := projectMap[projectID]
		if !exists {
			projectMap[projectID] = &dto.ProjectGetAllDTO{
				ID:            projectID,
				Title:         row.Title,
				Status:        row.Status,
				Summary:       row.Summary,
				ImageURL:      row.ImageURL,
				RepositoryURL: row.RepositoryURL,
				PublishedAt:   row.PublishedAt,
				Technologies:  []dto.ProjectTechnologyDTO{},
			}
		}

		projectMap[projectID].Technologies = append(projectMap[projectID].Technologies, dto.ProjectTechnologyDTO{
			ID:   row.TechID,
			Name: row.TechName,
		})
	}

	utils.PrintJSON(projectMap)

	// Convert map to slice
	var result []dto.ProjectGetAllDTO
	for _, project := range projectMap {
		result = append(result, *project)
	}

	return result, nil
}

func GetProject(id uint) (dto.ProjectSingleResponse, error) {
	db, _ := config.DB.DB()

	query := `
		SELECT 
			p.id,
			p.title,
			p.description,
			p.image_url,
			p.image_file_name,
			p.repository_url,
			p.summary,
			p.status,
			p.statistic_id,
			p.published_at,
			p.created_at,
			s.id AS statistic_id,
			s.likes AS statistic_likes,
			s.views AS statistic_views,
			s.type AS statistic_type,
			s.created_at AS statistic_created_at
		FROM projects p
		JOIN statistics s ON s.id = p.statistic_id
		WHERE p.id = ?
	`

	rows, err := db.Query(query, id)
	if err != nil {
		utils.LogError(err.Error(), query)
		return dto.ProjectSingleResponse{}, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var project models.Project
	if rows.Next() {
		if err := rows.Scan(
			&project.ID,
			&project.Title,
			&project.Description,
			&project.ImageURL,
			&project.ImageFileName,
			&project.RepositoryURL,
			&project.Summary,
			&project.Status,
			&project.StatisticId,
			&project.PublishedAt,
			&project.CreatedAt,
			&project.Statistic.ID,
			&project.Statistic.Likes,
			&project.Statistic.Views,
			&project.Statistic.Type,
			&project.Statistic.CreatedAt); err != nil {
			utils.LogWarning(err.Error(), query)
			return dto.ProjectSingleResponse{}, err
		}
	}

	var publishedAt *string = nil
	if project.PublishedAt != nil {
		formattedPublishedAt := project.PublishedAt.Format("2006-01-02")
		publishedAt = &formattedPublishedAt
	}

	statistics := dto.StatisticResponse{
		ID:        project.Statistic.ID,
		Likes:     project.Statistic.Likes,
		Views:     project.Statistic.Views,
		Type:      project.Statistic.Type,
		CreatedAt: project.Statistic.CreatedAt.Format("2006-01-02"),
	}

	return dto.ProjectSingleResponse{
		ID:            project.ID,
		Title:         project.Title,
		Description:   project.Description,
		ImageURL:      *project.ImageURL,
		ImageFileName: *project.ImageFileName,
		RepositoryURL: project.RepositoryURL,
		Summary:       project.Summary,
		Status:        project.Status,
		StatisticId:   project.StatisticId,
		PublishedAt:   publishedAt,
		CreatedAt:     project.CreatedAt.Format("2006-01-02"),
		Statistic:     statistics,
	}, nil
}

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

func InsertStatisticProject() (response models.Statistic, err error) {
	var likes int = 0
	var views int = 0
	resStatistic := models.Statistic{
		Likes: &likes,
		Views: &views,
		Type:  "Project",
	}

	if err := config.DB.Create(&resStatistic).Error; err != nil {
		return response, err
	}

	return resStatistic, nil
}

func InsertTechnologyProject(technologyIds []string, projectId int) (err error) {
	// INSERT to TABLE PROJECT TECHONOLOGIES
	var technologies []models.ProjectTechnology

	for _, technology_id := range technologyIds {
		TechnologyId, _ := strconv.Atoi(technology_id)
		technologies = append(technologies, models.ProjectTechnology{
			ProjectID:    projectId,
			TechnologyID: TechnologyId,
		})
	}

	if err := config.DB.Table("project_technologies").Create(&technologies).Error; err != nil {
		return err
	}

	return nil
}

func UpdateImagesProject(content_images []string, projectId int) (err error) {
	// UPDATE to TABLE PROJECT CONTENT IMAGES
	err = config.DB.Model(&models.ProjectContentImage{}).
		Where("image_url IN ?", content_images).
		Update("project_id", projectId).Error

	if err != nil {
		return err
	}

	return nil
}

func CreateProject(req dto.CreateProjectRequest, c *gin.Context) (result dto.ProjectCreateResponse, statusCode int, errFiels []utils.FieldError, err error) {
	tx := config.DB.Begin()

	// INSERT statistic
	resStatistic, err := InsertStatisticProject()
	if err != nil {
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
		StatisticId:   resStatistic.ID,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		tx.Rollback()
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return result, http.StatusInternalServerError, nil, err
	}

	// INSERT to TABLE PROJECT TECHONOLOGIES
	err = InsertTechnologyProject(req.TechnologyIds, int(data.ID))

	if err != nil {
		tx.Rollback()
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return result, http.StatusInternalServerError, nil, err
	}

	// UPDATE to TABLE PROJECT CONTENT IMAGES
	err = UpdateImagesProject(req.ContentImages, int(data.ID))

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
