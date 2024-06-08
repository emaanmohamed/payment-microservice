package Utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateAndRespond(c *gin.Context, request interface{}) bool {
	validate := validator.New()
	err := validate.Struct(request)
	if err == nil {
		return true
	}

	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = getErrorMessage(err)
	}

	c.JSON(400, gin.H{"errors": errors})
	return false
}

func getErrorMessage(err validator.FieldError) string {
	switch err.Field() {
	case "Gateway":
		return "Gateway is required"
	case "Amount":
		return "Amount is required and must be greater than 0"
	case "Currency":
		return "Currency is required and must be 3 characters long"
	case "Description":
		return "Description is required"
	case "Status":
		return "Status is required"
	default:
		return "Invalid request"
	}
}
