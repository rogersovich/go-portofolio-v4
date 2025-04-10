package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type SQLFilter struct {
	Column string
	Value  interface{}
	Op     string // e.g. "LIKE", "=", "IS NULL"
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

// BindJSON parses and validates JSON, with custom error responses for type issues
func BindJSON(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		var unmarshalTypeError *json.UnmarshalTypeError
		if errors.As(err, &unmarshalTypeError) {
			Error(c, http.StatusBadRequest,
				fmt.Sprintf("Invalid type for field '%s'. Expected %s",
					unmarshalTypeError.Field, unmarshalTypeError.Type.String()))
			return false
		}

		var syntaxError *json.SyntaxError
		if errors.As(err, &syntaxError) {
			Error(c, http.StatusBadRequest, "Malformed JSON request")
			return false
		}

		Error(c, http.StatusBadRequest, "Invalid request payload")
		return false
	}

	return true
}

func BuildSQLFilters(filters []SQLFilter) (conditions []string, args []interface{}) {
	for _, f := range filters {
		// fmt.Println("value:", f.Value, f.Column, f.Op)
		if f.Value == nil || f.Value == "" {
			continue // Skip empty values
		}

		switch f.Op {
		case "LIKE":
			conditions = append(conditions, f.Column+" LIKE ?")
			args = append(args, "%"+f.Value.(string)+"%")
		case "=":
			conditions = append(conditions, f.Column+" = ?")
			args = append(args, f.Value)
		case "IS NULL", "IS NOT NULL":
			conditions = append(conditions, f.Column+" "+f.Op)
		}
	}
	return
}

// BuildWhereClause combines all conditions into a SQL WHERE clause
func BuildWhereClause(conditions []string) string {
	if len(conditions) == 0 {
		return ""
	}
	return " WHERE " + strings.Join(conditions, " AND ")
}

// BuildOrderAndPagination returns the ORDER BY, LIMIT, and OFFSET clause
func BuildOrderAndPagination(order, sort string, page, limit int) string {
	// Default order column
	if order == "" {
		order = "id"
	}

	// Default sort direction
	sort = strings.ToUpper(sort)
	if sort != "DESC" {
		sort = "ASC"
	}

	// Default limit and page
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	return fmt.Sprintf(" ORDER BY %s %s LIMIT %d OFFSET %d", order, sort, limit, offset)
}
