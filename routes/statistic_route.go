package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/controllers"
)

func RegisterStatisticRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		statistic := api.Group("/statistics")
		{
			statistic.GET("", controllers.GetAllStatistics)
			statistic.GET("/:id", controllers.GetStatistic)
			statistic.POST("/store", controllers.CreateStatistic)
			statistic.POST("/update", controllers.UpdateStatistic)
			statistic.POST("/delete", controllers.DeleteStatistic)
		}
	}
}
