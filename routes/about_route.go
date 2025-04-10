package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/controllers"
)

func RegisterAboutRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		tech := api.Group("/abouts")
		{
			tech.GET("", controllers.GetAllAbouts)
			// tech.GET("/:id", controllers.GetAbout)
			// tech.POST("/store", controllers.CreateTechnology)
			// tech.POST("/update/:id", controllers.UpdateTechnology)
			// tech.POST("/delete", controllers.DeleteTechnology)
		}
	}
}
