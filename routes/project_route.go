package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/controllers"
)

func RegisterProjectRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		project := api.Group("/projects")
		{
			project.GET("", controllers.GetAllProjects)
			project.GET("/all-with-split", controllers.GetAllWithSplitQuery)
			project.GET("/:id", controllers.GetProject)
			project.POST("/store", controllers.CreateProject)
			// project.POST("/update", controllers.UpdateProject)
			// project.POST("/delete", controllers.DeleteProject)
		}
	}
}
