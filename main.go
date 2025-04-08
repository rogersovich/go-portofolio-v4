package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rogersovich/go-portofolio-v4/config"
)

func main() {
	// Init Viper Config
	config.InitConfig()

	// Init DB
	config.ConnectDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Welcome to %s!", config.Config.Name),
		})
	})

	r.Run(fmt.Sprintf(":%d", config.Config.Port))
}
