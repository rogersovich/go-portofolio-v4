package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Name  string
	Port  int
	Debug bool
}

var Config AppConfig

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}

	Config = AppConfig{
		Name:  viper.GetString("app.name"),
		Port:  viper.GetInt("app.port"),
		Debug: viper.GetBool("app.debug"),
	}
}
