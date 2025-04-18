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
			tech.GET("/:id", controllers.GetTechnology)
			tech.POST("/store", controllers.CreateTechnology)
			tech.POST("/update", controllers.UpdateTechnology)
			tech.POST("/delete", controllers.DeleteTechnology)
		}
	}
}
