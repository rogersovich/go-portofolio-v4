package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/config"
)

func main() {
	// Init Viper Config
	config.InitConfig()

	// Set Gin mode
	if config.Config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hellow world roger",
		})
	})

	port := fmt.Sprintf(":%d", config.Config.Port)
	r.Run(port) // listen and serve
}
