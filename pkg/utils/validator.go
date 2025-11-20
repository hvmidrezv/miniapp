package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			switch e.Tag() {
			case "required":
				return e.Field() + " is required"
			case "email":
				return e.Field() + " must be a valid email"
			case "oneof":
				return e.Field() + " must be one of: " + e.Param()
			default:
				return e.Field() + " is invalid"
			}
		}
	}
	return err.Error()
}

func GetPaginationParams(c *gin.Context) (page int, pageSize int) {
	page = 1
	pageSize = 10

	if p := c.Query("page"); p != "" {
		if parsedPage := parseIntOrDefault(p, 1); parsedPage > 0 {
			page = parsedPage
		}
	}

	if ps := c.Query("page_size"); ps != "" {
		if parsedSize := parseIntOrDefault(ps, 10); parsedSize > 0 && parsedSize <= 100 {
			pageSize = parsedSize
		}
	}

	return page, pageSize
}

func parseIntOrDefault(s string, defaultValue int) int {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	if err != nil {
		return defaultValue
	}
	return result
}
