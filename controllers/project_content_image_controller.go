package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/dto"
	"github.com/rogersovich/go-portofolio-v4/services"
	"github.com/rogersovich/go-portofolio-v4/utils"
)

func CreateProjectContentImage(c *gin.Context) {
	// Get text fields
	project_id := c.PostForm("project_id")
	is_used := c.PostForm("is_used")
	image_file := c.PostForm("image_file")

	// Validate the struct using validator
	req := dto.CreateProjectContentImageRequest{
		ProjectId: &project_id,
		IsUsed:    is_used,
		ImageFile: &image_file,
	}

	if verr := utils.ValidateRequest(&req); verr != nil {
		utils.ErrorValidation(c, http.StatusBadRequest, "Validation Error", verr)
		return
	}

	response, statusCode, errField, err := services.CreateProjectContentImage(req, c)
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

func GetProjectContentImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	tech, err := services.GetProjectContentImage(id)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "Success fetched data", tech)
}
