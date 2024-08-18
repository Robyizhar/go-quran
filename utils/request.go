package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPaginationParameters extracts pagination parameters from the context.
func GetPaginationParameters(c *gin.Context) (page int, limit int, orderBy string, orderType string, offset int) {
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ = strconv.Atoi(c.DefaultQuery("limit", "10"))
	orderBy = c.DefaultQuery("orderBy", "id")
	orderType = c.DefaultQuery("orderType", "asc")
	offset = (page - 1) * limit

	// Validate orderType
	if orderType != "asc" && orderType != "ASC" && orderType != "desc" && orderType != "DESC" {
		orderType = "asc"
	}

	return
}

func ValidateOrderBy(orderBy string, validOrderBys map[string]bool) string {
	if !validOrderBys[orderBy] {
		return "id" // Default to "id" if the provided orderBy is not valid
	}
	return orderBy
}
