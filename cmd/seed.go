package main

import (
	"github.com/rogersovich/go-portofolio-v4/config"
	"github.com/rogersovich/go-portofolio-v4/seeds"
)

func runSeed() {
	// Init Config
	config.InitConfigForTest()

	// Init DB
	config.ConnectDB()

	// Init Seed
	seeds.SeedAll()
}
