package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/dto"
	"github.com/rogersovich/go-portofolio-v4/services"
	"github.com/rogersovich/go-portofolio-v4/utils"
)

func GetAllAbouts(c *gin.Context) {
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

	params := dto.AboutQueryParams{
		Sort:        c.DefaultQuery("sort", "ASC"),
		Order:       c.DefaultQuery("order", "id"),
		Title:       c.Query("title"),
		Description: c.Query("description"),
		IsDelete:    c.Query("is_delete"), // expects "Y" or "N"
		CreatedFrom: c.Query("created_from"),
		CreatedTo:   c.Query("created_to"),
		Page:        page,
		Limit:       limit,
	}

	response, err := services.GetAllAbouts(params)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "Success fetched data", response)
}

func GetAbout(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	tech, err := services.GetAbout(id)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "Success fetched data", tech)
}

func CreateAbout(c *gin.Context) {
	// Get text fields
	title := c.PostForm("title")
	description := c.PostForm("description")
	avatarFile := c.PostForm("avatar_file")

	// Validate the struct using validator
	req := dto.CreateAboutRequest{
		Title:           title,
		DescriptionHTML: description,
		AvatarFile:      avatarFile,
	}

	if verr := utils.ValidateRequest(&req); verr != nil {
		utils.ErrorValidation(c, http.StatusBadRequest, "Validation Error", verr)
		return
	}

	response, statusCode, errField, err := services.CreateAbout(req, c)
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

func UpdateAbout(c *gin.Context) {
	// Get text fields
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	title := c.PostForm("title")
	description := c.PostForm("description")

	// Validate the struct using validator
	req := dto.UpdateAboutRequest{
		Id:              id,
		Title:           title,
		DescriptionHTML: description,
	}

	if verr := utils.ValidateRequest(&req); verr != nil {
		utils.ErrorValidation(c, http.StatusBadRequest, "Validation Error", verr)
		return
	}

	//Update Database
	response, statusCode, errField, err := services.UpdateAbout(req, id, c)
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

func DeleteAbout(c *gin.Context) {
	var req dto.DeleteAboutRequest

	if !utils.ValidateStruct(c, &req, c.ShouldBindJSON(&req)) {
		return
	}

	id := req.ID

	// Delete data
	statusCode, err := services.DeleteAbout(id, c)
	if err != nil {
		utils.Error(c, statusCode, err.Error())
		return
	}

	utils.Success(c, "Success to delete data", nil)
}
