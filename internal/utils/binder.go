package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValidate[T any](c *gin.Context, req *T) bool {
	// parsing JSON
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid request payload: " + err.Error(),
		})
		return false
	}

	// validasi
	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errMap := make(map[string]string)

		for _, fe := range validationErrors {
			errMap[fe.Field()] = validationErrorToText(fe)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errors":  errMap,
		})
		return false
	}

	return true
}

func validationErrorToText(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return fe.Field() + " must be a valid email"
	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters"
	case "max":
		return fe.Field() + " must be at most " + fe.Param() + " characters"
	}
	return fe.Field() + " is invalid"
}
