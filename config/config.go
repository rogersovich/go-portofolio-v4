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

	// Load .env (default)
	viper.SetConfigFile(".env")
	_ = viper.MergeInConfig()

	// Load .env.production if exists (for deployment)
	viper.SetConfigFile(".env.production")
	_ = viper.MergeInConfig()

	// Load environment variables (overrides everything)
	viper.AutomaticEnv()

	Config = AppConfig{
		Name: viper.GetString("app.name"),
		Port: viper.GetInt("app.port"),
		Database: DatabaseConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetInt("database.port"),
			User:     viper.GetString("database.user"),
			Password: viper.GetString("database.password"),
			Name:     viper.GetString("database.name"),
		},
	}
}
