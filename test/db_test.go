package test

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/rogersovich/go-portofolio-v4/config"
)

func TestDBConnect(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		_ = godotenv.Load("../.env") // fallback if not found in local dir
	}

	config.InitConfigForTest()
}
