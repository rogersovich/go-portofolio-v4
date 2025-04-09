package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/controllers"
)

func RegisterTechnologyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		tech := api.Group("/technologies")
		{
			tech.GET("", controllers.GetAllTechnologies)
			// future: tech.POST, tech.PUT, etc.
		}
	}
}
