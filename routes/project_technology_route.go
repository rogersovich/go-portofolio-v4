package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/controllers"
)

func RegisterProjectTechnologyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		project_technology := api.Group("/project-technologies")
		{
			// project_technology.GET("", controllers.GetAllProjectTechnologies)
			project_technology.GET("/:id", controllers.GetProjectTechnology)
			project_technology.POST("/store", controllers.CreateProjectTechnology)
			project_technology.POST("/update", controllers.UpdateProjectTechnology)
			project_technology.POST("/delete", controllers.DeleteProjectTechnology)
		}
	}
}
