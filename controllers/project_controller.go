package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/dto"
	"github.com/rogersovich/go-portofolio-v4/services"
	"github.com/rogersovich/go-portofolio-v4/utils"
)

func CreateProject(c *gin.Context) {
	// Get text fields
	title := c.PostForm("title")
	description := c.PostForm("description")
	image_file := c.PostForm("image_file")
	repository_url := c.PostForm("repository_url")
	summary := c.PostForm("summary")
	status := c.PostForm("status")             // "Published", "Unpublished", "Deleted"
	is_published := c.PostForm("is_published") // "Y", "N"
	technology_ids := c.PostFormArray("technology_ids[]")
	content_images := c.PostFormArray("content_images[]")

	technology_ids_validated, err := utils.ValidateFormArrayNotEmpty(technology_ids, "technology_ids", true)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	content_images_validated, err := utils.ValidateFormArrayNotEmpty(content_images, "content_images", false)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// Validate the struct using validator
	req := dto.CreateProjectRequest{
		Title:         title,
		Description:   description,
		ImageFile:     image_file,
		RepositoryURL: repository_url,
		Summary:       summary,
		Status:        status,
		IsPublihed:    is_published,
		TechnologyIds: technology_ids_validated,
		ContentImages: content_images_validated,
	}

	if verr := utils.ValidateRequest(&req); verr != nil {
		utils.ErrorValidation(c, http.StatusBadRequest, "Validation Error", verr)
		return
	}

	err = services.CheckProjectTechnology(req.TechnologyIds)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// response, statusCode, errField, err := services.CreateProject(req, c)
	// if err != nil {
	// 	if statusCode == http.StatusBadRequest {
	// 		utils.ErrorValidation(c, http.StatusBadRequest, err.Error(), errField)
	// 	} else {
	// 		utils.Error(c, http.StatusInternalServerError, err.Error())
	// 	}
	// 	return
	// }

	utils.Success(c, "Success to create data", req)
}
