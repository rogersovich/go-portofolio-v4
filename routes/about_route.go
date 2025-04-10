package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/controllers"
)

func RegisterAboutRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		about := api.Group("/abouts")
		{
			about.GET("", controllers.GetAllAbouts)
			// about.GET("/:id", controllers.GetAbout)
			about.POST("/store", controllers.CreateAbout)
			// about.POST("/update/:id", controllers.UpdateAbout)
			// about.POST("/delete", controllers.DeleteAbout)
		}
	}
}
