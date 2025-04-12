package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/controllers"
)

func RegisterAuthorRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		author := api.Group("/authors")
		{
			author.GET("", controllers.GetAllAuthors)
			author.GET("/:id", controllers.GetAuthor)
			author.POST("/store", controllers.CreateAuthor)
			author.POST("/update", controllers.UpdateAuthor)
			author.POST("/delete", controllers.DeleteAuthor)
		}
	}
}
