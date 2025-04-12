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

var folderNameAuthor = "author"

func GetAllAuthors(params dto.AuthorQueryParams) ([]dto.AuthorResponse, error) {
	db, _ := config.DB.DB()

	var (
		conditions []string
		args       []interface{}
	)

	query := `SELECT id, name, avatar_url, avatar_file_name, created_at FROM authors`

	// 🔍 Filters

	filters := []utils.SQLFilter{
		{Column: "name", Value: params.Name, Op: "LIKE"},
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

	var authors []models.Author

	for rows.Next() {
		var rowAuthor models.Author
		if err := rows.Scan(&rowAuthor.ID, &rowAuthor.Name, &rowAuthor.AvatarUrl, &rowAuthor.AvatarFileName, &rowAuthor.CreatedAt); err != nil {
			utils.LogWarning(err.Error(), query)
			return nil, err
		}
		authors = append(authors, rowAuthor)
	}

	var responses []dto.AuthorResponse
	for _, rowAuthor := range authors {
		responses = append(responses, dto.AuthorResponse{
			ID:             rowAuthor.ID,
			Name:           rowAuthor.Name,
			AvatarURL:      rowAuthor.AvatarUrl,
			AvatarFileName: rowAuthor.AvatarFileName,
			CreatedAt:      rowAuthor.CreatedAt.Format("2006-01-02"),
		})
	}

	return responses, nil
}

func GetAuthor(id int) (dto.AuthorSingleResponse, error) {
	var response models.Author
	if err := config.DB.First(&response, id).Error; err != nil {
		return dto.AuthorSingleResponse{}, err
	}

	return dto.AuthorSingleResponse{
		ID:             response.ID,
		Name:           response.Name,
		AvatarURL:      response.AvatarUrl,
		AvatarFileName: response.AvatarFileName,
		CreatedAt:      response.CreatedAt.Format("2006-01-02"),
	}, nil
}

func CreateAuthor(req dto.CreateAuthorRequest, c *gin.Context) (result dto.AuthorSingleResponse, statusCode int, errFiels []utils.FieldError, err error) {
	// Upload avatar_file
	avatarData, avatarErrs, avatarUploadErr := uploadService.HandleUploadedFile(
		c,
		"avatar_file",
		folderNameAuthor,
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

	data := models.Author{
		Name:           req.Name,
		AvatarUrl:      avatarData.FileURL,
		AvatarFileName: &avatarData.FileName,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		return result, http.StatusInternalServerError, nil, err
	}

	result = dto.AuthorSingleResponse{
		ID:             data.ID,
		Name:           data.Name,
		AvatarURL:      data.AvatarUrl,
		AvatarFileName: data.AvatarFileName,
		CreatedAt:      data.CreatedAt.Format("2006-01-02"),
	}

	return result, http.StatusOK, nil, nil
}

func UpdateAuthor(req dto.UpdateAuthorRequest, id int, c *gin.Context) (result dto.AuthorUpdateResponse, statusCode int, errFiels []utils.FieldError, err error) {
	// 1. Fetch existing Author data
	resAuthor, err := GetAuthor(id)
	if err != nil {
		return result, http.StatusNotFound, nil, err
	}

	// set oldPath
	oldPath := ""
	if resAuthor.AvatarFileName != nil {
		oldPath = *resAuthor.AvatarFileName
	}

	// 2. Get new file (if uploaded)
	_, err = c.FormFile("avatar_file")
	var newFileURL string
	var newFileName string

	if err == nil {
		// Upload avatar_file
		avatarData, avatarErrs, avatarUploadErr := uploadService.HandleUploadedFile(
			c,
			"avatar_file",
			folderNameAuthor,
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
		newFileURL = resAuthor.AvatarURL // keep existing if not updated
		newFileName = *resAuthor.AvatarFileName
	}

	data := models.Author{
		Name:           req.Name,
		AvatarUrl:      newFileURL,
		AvatarFileName: &newFileName,
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

	return dto.AuthorUpdateResponse{
		Name:           data.Name,
		AvatarURL:      data.AvatarUrl,
		AvatarFileName: data.AvatarFileName,
	}, http.StatusOK, nil, nil
}

func DeleteAuthor(id int, c *gin.Context) (statusCode int, err error) {
	// 1. Fetch existing Author data
	_, err = GetAuthor(id)
	if err != nil {
		return http.StatusNotFound, err
	}

	// 2. Optional: Delete old file from MinIO
	// oldPath := resAuthor.AvatarFileName
	// err = uploadService.DeleteFromMinio(c.Request.Context(), oldPath)
	// if err != nil {
	// 	return http.StatusInternalServerError, err
	// }

	// 3. Delete data
	if err := config.DB.Delete(&models.Author{}, id).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
