package config

import (
	"github.com/spf13/viper"
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
	// Load config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // root path
	_ = viper.ReadInConfig() // optional, don't panic if missing

	// Load .env
	envViper := viper.New()
	envViper.SetConfigFile(".env")
	if err := envViper.ReadInConfig(); err == nil {
		_ = viper.MergeConfigMap(envViper.AllSettings())
	}

	// Load .env.production if it exists (optional override)
	envProd := viper.New()
	envProd.SetConfigFile(".env.production")
	if err := envProd.ReadInConfig(); err == nil {
		_ = viper.MergeConfigMap(envProd.AllSettings())
	}

	// Bind environment variables
	viper.AutomaticEnv()
	_ = viper.BindEnv("DB_USER")
	_ = viper.BindEnv("DB_PASS")

	Config = AppConfig{
		Name: viper.GetString("app.name"),
		Port: viper.GetInt("app.port"),
		Database: DatabaseConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetInt("database.port"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASS"),
			Name:     viper.GetString("database.name"),
		},
	}
}
