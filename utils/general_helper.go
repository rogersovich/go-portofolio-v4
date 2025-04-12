package utils

import (
	"os"
	"strings"
)

func GetIsProduction() bool {
	env := strings.ToLower(os.Getenv("APP_ENV"))
	return env == "production"
}

func GetProtocol() string {
	isProduction := GetIsProduction()
	if isProduction {
		return "https"
	}

	return "http" // default development
}

func GetEnv(key string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return ""
}

func StringOrDefault(s *string, def string) *string {
	if s == nil {
		return &def
	}
	return s
}
