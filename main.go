package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/config"
	"github.com/rogersovich/go-portofolio-v4/routes"
	"github.com/rogersovich/go-portofolio-v4/utils"
)

func main() {
	utils.InitLogger()
	config.InitConfig()
	config.ConnectDB()

	r := gin.Default()

	// Init Routes
	routes.RegisterTechnologyRoutes(r)
	routes.RegisterAboutRoutes(r)
	routes.RegisterAuthorRoutes(r)
	routes.RegisterProjectFeatureRoutes(r)

	r.NoRoute(func(c *gin.Context) {
		utils.Error(c, http.StatusNotFound, "Route not found")
	})

	r.Run(fmt.Sprintf(":%d", config.Config.Port))
}
