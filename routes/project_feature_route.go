package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/controllers"
)

func RegisterProjectFeatureRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		project_feature := api.Group("/project_features")
		{
			project_feature.GET("", controllers.GetAllProjectFeatures)
			project_feature.GET("/:id", controllers.GetProjectFeature)
			project_feature.POST("/store", controllers.CreateProjectFeature)
			project_feature.POST("/update", controllers.UpdateProjectFeature)
			project_feature.POST("/delete", controllers.DeleteProjectFeature)
		}
	}
}
