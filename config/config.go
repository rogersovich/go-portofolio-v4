package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Name     string
	Port     int
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

var Config AppConfig

func InitConfig() {
	// ✅ Load .env (priority)
	if err := godotenv.Load(".env"); err != nil {
		log.Println("⚠️ .env not found, trying .env.production...")
		_ = godotenv.Load(".env.production")
	}

	// 🌍 Load vars from environment
	appPort, _ := strconv.Atoi(getEnv("APP_PORT"))
	dbPort, _ := strconv.Atoi(getEnv("DB_PORT"))

	Config = AppConfig{
		Name: getEnv("APP_NAME"),
		Port: appPort,
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST"),
			Port:     dbPort,
			User:     getEnv("DB_USER"),
			Password: getEnv("DB_PASSWORD"),
			Name:     getEnv("DB_NAME"),
		},
	}

	// Optional debug print
	fmt.Printf("✅ Loaded config: %+v\n", Config)
}

func InitConfigForTest() {
	// ✅ Load .env (priority)
	if err := godotenv.Load(".env"); err != nil {
		_ = godotenv.Load(".env.production")
	}

	// 🌍 Load vars from environment
	appPort, _ := strconv.Atoi(getEnv("APP_PORT"))
	dbPort, _ := strconv.Atoi(getEnv("DB_PORT"))

	Config = AppConfig{
		Name: getEnv("APP_NAME"),
		Port: appPort,
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST"),
			Port:     dbPort,
			User:     getEnv("DB_USER"),
			Password: getEnv("DB_PASSWORD"),
			Name:     getEnv("DB_NAME"),
		},
	}
}

func getEnv(key string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return ""
}
