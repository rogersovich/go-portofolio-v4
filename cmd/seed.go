package main

import (
	"github.com/rogersovich/go-portofolio-v4/config"
	"github.com/rogersovich/go-portofolio-v4/seeds"
)

func main() {
	// Init Config
	config.InitConfig()

	config.ConnectDB()

	// Init Seed
	seeds.SeedAll()
}
