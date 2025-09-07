package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Success response
func RespondSuccess(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error response
func RespondError(c *gin.Context, status int, message string) {
	c.JSON(status, Response{
		Success: false,
		Error:   message,
	})
}

// Shortcut common errors
func BadRequest(c *gin.Context, msg string) {
	RespondError(c, http.StatusBadRequest, msg)
}

func Unauthorized(c *gin.Context, msg string) {
	RespondError(c, http.StatusUnauthorized, msg)
}

func InternalError(c *gin.Context, msg string) {
	RespondError(c, http.StatusInternalServerError, msg)
}