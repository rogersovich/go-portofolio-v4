package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/dto"
	"github.com/rogersovich/go-portofolio-v4/services"
	"github.com/rogersovich/go-portofolio-v4/utils"
)

func GetAllTechnologies(c *gin.Context) {
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

	params := dto.TechnologyQueryParams{
		Sort:        c.DefaultQuery("sort", "ASC"),
		Order:       c.DefaultQuery("order", "id"),
		FilterName:  c.Query("name"),
		FilterDesc:  c.Query("description"),
		IsMajor:     c.Query("is_major"),  // expects "Y" or "N"
		IsDelete:    c.Query("is_delete"), // expects "Y" or "N"
		CreatedFrom: c.Query("created_from"),
		CreatedTo:   c.Query("created_to"),
		Page:        page,
		Limit:       limit,
	}

	technologies, err := services.GetAllTechnologies(params)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "Technologies fetched successfully", technologies)
}

func GetTechnology(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid technology ID")
		return
	}

	tech, err := services.GetTechnology(id)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "Technology fetched successfully", tech)
}

func CreateTechnology(c *gin.Context) {
	// Get text fields
	name := c.PostForm("name")
	description := c.PostForm("description")
	isMajor := c.PostForm("is_major")
	logoFile := c.PostForm("logo_file")

	// Validate the struct using validator
	req := dto.CreateTechnologyRequest{
		Name:            name,
		DescriptionHTML: description,
		LogoFile:        logoFile,
		IsMajor:         isMajor,
	}

	if verr := utils.ValidateRequest(&req); verr != nil {
		utils.ErrorValidation(c, http.StatusBadRequest, "Validation Error", verr)
		return
	}

	response, statusCode, errField, err := services.CreateTechnology(req, c)
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

func UpdateTechnology(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	// Get text fields
	name := c.PostForm("name")
	description := c.PostForm("description")
	isMajor := c.PostForm("is_major")

	// Validate the struct using validator
	req := dto.UpdateTechnologyRequest{
		Id:              id,
		Name:            name,
		DescriptionHTML: description,
		IsMajor:         isMajor,
	}

	if verr := utils.ValidateRequest(&req); verr != nil {
		utils.ErrorValidation(c, http.StatusBadRequest, "Validation Error", verr)
		return
	}

	response, statusCode, errField, err := services.UpdateTechnology(req, id, c)
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

func DeleteTechnology(c *gin.Context) {
	var req dto.DeleteTechnologyRequest

	if !utils.BindJSON(c, &req) || !utils.ValidateStruct(c, &req, nil) {
		return
	}

	id := req.ID

	tech, err := services.DeleteTechnology(id)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "Technology deleted successfully", tech)
}
