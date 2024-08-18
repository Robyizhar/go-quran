package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, data interface{}, params ...interface{}) {
	statusCode := http.StatusOK
	message := "Request successful"
	var meta interface{} = nil

	if len(params) > 0 {
		if code, ok := params[1].(int); ok {
			statusCode = code
		}
	}

	if len(params) > 1 {
		if msg, ok := params[1].(string); ok {
			message = msg
		}
	}

	if len(params) > 2 {
		meta = params[2]
	}

	response := gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	}

	if meta != nil {
		response["meta"] = meta
	}

	c.JSON(statusCode, response)
}

func ErrorResponse(c *gin.Context, params ...interface{}) {
	message := "An error occurred"
	statusCode := http.StatusInternalServerError

	if len(params) > 0 {
		if code, ok := params[0].(int); ok {
			statusCode = code
		}
	}

	if len(params) > 1 {
		if msg, ok := params[1].(string); ok {
			message = msg
		}
	}

	c.JSON(statusCode, gin.H{
		"status":  "error",
		"message": message,
		"data":    nil,
	})
}
