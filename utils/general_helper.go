package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
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

// Converts a slice of string (like from c.PostFormArray) to a slice of int
func ValidateFormArrayToIntSlice(strs []string, field string, is_required bool) ([]int, error) {
	var result []int

	for _, s := range strs {
		if s == "" {
			continue
		}

		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, errors.New(field + "invalid integer value in array: " + s)
		}
		result = append(result, i)
	}

	if is_required && len(result) == 0 {
		return nil, errors.New(field + " array must not be empty")
	}

	// Ensure empty slice (not nil) is returned if not required
	if !is_required && len(result) == 0 {
		return []int{}, nil
	}

	return result, nil
}

// Validates that a string array is not empty and doesn't contain only empty strings
func ValidateFormArrayString(strs []string, field string, is_required bool) ([]string, error) {
	var result []string

	for _, s := range strs {
		if s != "" {
			result = append(result, s)
		}
	}

	if is_required && len(result) == 0 {
		return nil, errors.New(field + "array must not be empty")
	}

	// Ensure empty slice (not nil) is returned if not required
	if !is_required && len(result) == 0 {
		return []string{}, nil
	}

	return result, nil
}

func BoolToYN(val bool) string {
	if val {
		return "Y"
	}
	return "N"
}

func StringBoolToYN(val string) string {
	if val == "1" {
		return "Y"
	}
	return "N"
}

func PrintJSON(v any) {
	jsonBytes, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(jsonBytes))
}
