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

func GetAbout(id int) (dto.AboutSingleResponse, error) {
	var about models.About
	if err := config.DB.First(&about, id).Error; err != nil {
		return dto.AboutSingleResponse{}, err
	}

	return dto.AboutSingleResponse{
		ID:              about.ID,
		Title:           about.Title,
		DescriptionHTML: about.DescriptionHTML,
		AvatarURL:       about.AvatarUrl,
		AvatarFileName:  about.AvatarFileName,
		CreatedAt:       about.CreatedAt.Format("2006-01-02"),
	}, nil
}

func CreateAbout(req dto.CreateAboutRequest, c *gin.Context) (result dto.AboutSingleResponse, statusCode int, errFiels []utils.FieldError, err error) {
	// Upload avatar_file
	avatarData, avatarErrs, avatarUploadErr := uploadService.HandleUploadedFile(
		c,
		"avatar_file",
		"about",
		nil,         // use default allowed extensions
		2*1024*1024, // max 2MB
		nil,         // []string{"required", "extension", "size"}
	)

	if avatarErrs != nil {
		err = fmt.Errorf("invalid avatar_file")
		return result, http.StatusBadRequest, avatarErrs, err
	}

	if avatarUploadErr != nil {
		err = fmt.Errorf("failed to upload avatar_file")
		return result, http.StatusInternalServerError, avatarErrs, err
	}

	data := models.About{
		Title:           req.Title,
		DescriptionHTML: req.DescriptionHTML,
		AvatarUrl:       avatarData.FileURL,
		AvatarFileName:  avatarData.FileName,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		return result, http.StatusInternalServerError, nil, err
	}

	result = dto.AboutSingleResponse{
		ID:              data.ID,
		Title:           data.Title,
		DescriptionHTML: data.DescriptionHTML,
		AvatarURL:       data.AvatarUrl,
		AvatarFileName:  data.AvatarFileName,
		CreatedAt:       data.CreatedAt.Format("2006-01-02"),
	}

	return result, http.StatusOK, nil, nil
}

func UpdateAbout(req dto.UpdateAboutRequest, id int, c *gin.Context) (result dto.AboutUpdateResponse, statusCode int, errFiels []utils.FieldError, err error) {
	// 1. Fetch existing about data
	resAbout, err := GetAbout(id)
	if err != nil {
		return result, http.StatusNotFound, nil, err
	}

	// 2. Get new file (if uploaded)
	avatarFile, err := c.FormFile("avatar_file")
	var newFileURL string
	var newFileName string

	fmt.Printf("avatarFile: %v\n", avatarFile)

	if err == nil {
		// 3. Optional: Delete old file from MinIO
		oldPath := resAbout.AvatarFileName
		_ = uploadService.DeleteFromMinio(c.Request.Context(), oldPath) // ignore error or handle if needed

		// Upload avatar_file
		avatarData, avatarErrs, avatarUploadErr := uploadService.HandleUploadedFile(
			c,
			"avatar_file",
			"about",
			nil,                           // use default allowed extensions
			2*1024*1024,                   // max 2MB
			[]string{"extension", "size"}, // []string{"required", "extension", "size"}
		)

		if avatarErrs != nil {
			err = fmt.Errorf("invalid avatar_file")
			return result, http.StatusBadRequest, avatarErrs, err
		}

		if avatarUploadErr != nil {
			err = fmt.Errorf("failed to upload avatar_file")
			return result, http.StatusInternalServerError, avatarErrs, err
		}

		newFileURL = avatarData.FileURL
		newFileName = avatarData.FileName
	} else {
		newFileURL = resAbout.AvatarURL // keep existing if not updated
		newFileName = resAbout.AvatarFileName
	}

	data := models.About{
		Title:           req.Title,
		DescriptionHTML: req.DescriptionHTML,
		AvatarUrl:       newFileURL,
		AvatarFileName:  newFileName,
	}

	if err := config.DB.Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return result, http.StatusInternalServerError, nil, err
	}

	return dto.AboutUpdateResponse{
		Title:           data.Title,
		DescriptionHTML: data.DescriptionHTML,
		AvatarURL:       data.AvatarUrl,
		AvatarFileName:  data.AvatarFileName,
	}, http.StatusOK, nil, nil

}

func DeleteAbout(id int, c *gin.Context) (statusCode int, err error) {
	// 1. Fetch existing about data
	_, err = GetAbout(id)
	if err != nil {
		return http.StatusNotFound, err
	}

	// 2. Optional: Delete old file from MinIO
	// oldPath := resAbout.AvatarFileName
	// err = uploadService.DeleteFromMinio(c.Request.Context(), oldPath)
	// if err != nil {
	// 	return http.StatusInternalServerError, err
	// }

	// 3. Delete data
	if err := config.DB.Delete(&models.About{}, id).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
