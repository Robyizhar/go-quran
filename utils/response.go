package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SuccessResponse sends a successful response with data
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Request successful",
		"data":    data,
	})
}

// ErrorResponse sends an error response
func ErrorResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, gin.H{
		"status":  "error",
		"message": message,
		"data":    nil,
	})
}
