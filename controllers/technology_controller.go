package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/services"
	"github.com/rogersovich/go-portofolio-v4/utils"
)

type TechnologyResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Logo  string `json:"logo"`
	Major bool   `json:"is_major"`
}

func GetAllTechnologies(c *gin.Context) {
	technologies, err := services.GetAllTechnologies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch technologies"})
		return
	}

	var response []TechnologyResponse
	for _, tech := range technologies {
		response = append(response, TechnologyResponse{
			ID:    tech.ID,
			Name:  tech.Name,
			Logo:  tech.LogoURL,
			Major: tech.IsMajor,
		})
	}

	utils.Success(c, "Technologies fetched successfully", response)
}
