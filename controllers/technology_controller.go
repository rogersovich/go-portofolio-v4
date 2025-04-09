package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	params := services.TechnologyQueryParams{
		Sort:        c.DefaultQuery("sort", "ASC"),
		Order:       c.DefaultQuery("order", "id"),
		FilterName:  c.Query("filter_name"),
		FilterDesc:  c.Query("filter_desc"),
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
