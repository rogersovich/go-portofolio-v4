package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/dto"
	"github.com/rogersovich/go-portofolio-v4/services"
	"github.com/rogersovich/go-portofolio-v4/utils"
)

func GetAllProjectFeatures(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	params := dto.ProjectFeatureQueryParams{
		Sort:        c.DefaultQuery("sort", "ASC"),
		Order:       c.DefaultQuery("order", "id"),
		Description: c.Query("description"),
		IsDelete:    c.Query("is_delete"), // expects "Y" or "N"
		CreatedFrom: c.Query("created_from"),
		CreatedTo:   c.Query("created_to"),
		Page:        page,
		Limit:       limit,
	}

	response, err := services.GetAllProjectFeatures(params)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "Success fetched data", response)
}

func GetProjectFeature(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	tech, err := services.GetProjectFeature(id)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "Success fetched data", tech)
}

func CreateProjectFeature(c *gin.Context) {
	// Get text fields
	description := c.PostForm("description")
	imageFile := c.PostForm("image_file")

	// Validate the struct using validator
	req := dto.CreateProjectFeatureRequest{
		Description: description,
		ImageFile:   imageFile,
	}

	if verr := utils.ValidateRequest(&req); verr != nil {
		utils.ErrorValidation(c, http.StatusBadRequest, "Validation Error", verr)
		return
	}

	response, statusCode, errField, err := services.CreateProjectFeature(req, c)
	if err != nil {
		if statusCode == http.StatusBadRequest {
			utils.ErrorValidation(c, http.StatusBadRequest, err.Error(), errField)
		} else {
			utils.Error(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.Success(c, "Success to create data", response)
}

func UpdateProjectFeature(c *gin.Context) {
	// Get text fields
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	description := c.PostForm("description")

	// Validate the struct using validator
	req := dto.UpdateProjectFeatureRequest{
		Id:          id,
		Description: description,
	}

	if verr := utils.ValidateRequest(&req); verr != nil {
		utils.ErrorValidation(c, http.StatusBadRequest, "Validation Error", verr)
		return
	}

	//Update Database
	response, statusCode, errField, err := services.UpdateProjectFeature(req, id, c)
	if err != nil {
		if statusCode == http.StatusBadRequest {
			utils.ErrorValidation(c, http.StatusBadRequest, err.Error(), errField)
		} else {
			utils.Error(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.Success(c, "Success to update data", response)
}

func DeleteProjectFeature(c *gin.Context) {
	var req dto.DeleteProjectFeatureRequest

	if !utils.ValidateStruct(c, &req, c.ShouldBindJSON(&req)) {
		return
	}

	id := req.ID

	// Delete data
	statusCode, err := services.DeleteProjectFeature(id, c)
	if err != nil {
		utils.Error(c, statusCode, err.Error())
		return
	}

	utils.Success(c, "Success to delete data", nil)
}
