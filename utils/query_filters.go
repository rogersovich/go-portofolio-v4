package utils

import (
	"strings"
	"time"
)

// AddDateRangeFilter adds "field >= ?" and "field < ?" to WHERE clause if from/to are set.
// Automatically adds +1 day to 'to' to include full-day coverage.
func AddDateRangeFilter(field, from, to string, conditions *[]string, args *[]interface{}) {
	formatedLayout := "2006-01-02"

	if from != "" {
		fromTime, err := time.Parse(formatedLayout, strings.TrimSpace(from))
		if err == nil {
			*conditions = append(*conditions, field+" >= ?")
			*args = append(*args, fromTime.Format("2006-01-02 15:04:05"))
		}
	}

	if to != "" {
		toTime, err := time.Parse(formatedLayout, strings.TrimSpace(to))
		if err == nil {
			// add 1 day to make it exclusive
			toTime = toTime.AddDate(0, 0, 1)
			*conditions = append(*conditions, field+" < ?")
			*args = append(*args, toTime.Format("2006-01-02 15:04:05"))
		}
	}
}
