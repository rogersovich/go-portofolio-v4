package utils

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, message string, data interface{}) {

	if data == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": message,
			"data":    []interface{}{}, // default empty array
		})
		return
	}

	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	// If it's a slice and nil, return empty array
	if t.Kind() == reflect.Slice && v.IsNil() {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": message,
			"data":    []interface{}{}, // empty JSON array
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":  "error",
		"message": message,
	})
}

func ErrorValidation(c *gin.Context, code int, message string, errors interface{}) {
	c.JSON(code, gin.H{
		"status":  "error",
		"message": message,
		"errors":  errors,
	})
}
