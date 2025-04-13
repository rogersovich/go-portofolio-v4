package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/dto"
	"github.com/rogersovich/go-portofolio-v4/services"
	"github.com/rogersovich/go-portofolio-v4/utils"
)

func GetProjectTechnology(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	tech, err := services.GetProjectTechnology(id)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "Success fetched data", tech)
}

func CreateProjectTechnology(c *gin.Context) {
	var req dto.CreateProjectTechnologyRequest

	if !utils.ValidateStruct(c, &req, c.ShouldBindJSON(&req)) {
		return
	}

	response, err := services.CreateProjectTechnology(req)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "Success to create data", response)
}

func UpdateProjectTechnology(c *gin.Context) {
	var req dto.UpdateProjectTechnologyRequest

	if !utils.ValidateStruct(c, &req, c.ShouldBindJSON(&req)) {
		return
	}

	id := req.Id

	response, err := services.UpdateProjectTechnology(req, id)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "Success to create data", response)
}

func DeleteProjectTechnology(c *gin.Context) {
	var req dto.DeleteProjectTechnologyRequest

	if !utils.ValidateStruct(c, &req, c.ShouldBindJSON(&req)) {
		return
	}

	id := req.ID

	// Delete data
	statusCode, err := services.DeleteProjectTechnology(id, c)
	if err != nil {
		utils.Error(c, statusCode, err.Error())
		return
	}

	utils.Success(c, "Success to delete data", nil)
}
