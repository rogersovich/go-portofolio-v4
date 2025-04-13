package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/controllers"
)

func RegisterProjectContentImageRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		project_content_image := api.Group("/project-content-images")
		{
			// project_content_image.GET("", controllers.GetAllProjectContentImages)
			project_content_image.GET("/:id", controllers.GetProjectContentImage)
			project_content_image.POST("/store", controllers.CreateProjectContentImage)
			// project_content_image.POST("/update", controllers.UpdateProjectContentImage)
			// project_content_image.POST("/delete", controllers.DeleteProjectContentImage)
		}
	}
}
