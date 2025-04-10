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

	utils.Success(c, "About fetched successfully", response)
}

// func GetTechnology(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		utils.Error(c, http.StatusBadRequest, "Invalid technology ID")
// 		return
// 	}

// 	tech, err := services.GetTechnology(id)
// 	if err != nil {
// 		utils.Error(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	utils.Success(c, "Technology fetched successfully", tech)
// }

// func CreateTechnology(c *gin.Context) {
// 	var req dto.CreateTechnologyRequest

// 	if !utils.ValidateStruct(c, &req, c.ShouldBindJSON(&req)) {
// 		return // already responded with JSON errors
// 	}

// 	tech, err := services.CreateTechnology(req)
// 	if err != nil {
// 		utils.Error(c, http.StatusInternalServerError, "Failed to create data")
// 		return
// 	}

// 	utils.Success(c, "Technology created successfully", tech)
// }

// func UpdateTechnology(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		utils.Error(c, http.StatusBadRequest, "Invalid technology ID")
// 		return
// 	}

// 	var req dto.UpdateTechnologyRequest

// 	if !utils.ValidateStruct(c, &req, c.ShouldBindJSON(&req)) {
// 		return // already responded with JSON errors
// 	}

// 	tech, err := services.UpdateTechnology(req, id)
// 	if err != nil {
// 		utils.Error(c, http.StatusInternalServerError, "Failed to updated data")
// 		return
// 	}

// 	utils.Success(c, "Technology updated successfully", tech)
// }

// func DeleteTechnology(c *gin.Context) {
// 	var req dto.DeleteTechnologyRequest

// 	if !utils.BindJSON(c, &req) || !utils.ValidateStruct(c, &req, nil) {
// 		return
// 	}

// 	id := req.ID

// 	tech, err := services.DeleteTechnology(id)
// 	if err != nil {
// 		utils.Error(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	utils.Success(c, "Technology deleted successfully", tech)
// }
