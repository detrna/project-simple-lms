package shared

import (
	"errors"
	"fmt"
	"main/internal/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationError(c *gin.Context, logger pkg.Logger, err error) {
	var validationErrs validator.ValidationErrors

	if errors.As(err, &validationErrs) {
		errors := make(map[string]string)

		for _, e := range validationErrs {
			field := strings.ToLower(e.Field())

			switch e.Tag() {
			case "required":
				errors[field] = "is required"
			case "email":
				errors[field] = "must be a valid email"
			case "min":
				errors[field] = fmt.Sprintf("must be at least %s characters", e.Param())
			case "max":
				errors[field] = fmt.Sprintf("must be at most %s characters", e.Param())
			case "len":
				errors[field] = fmt.Sprintf("must be exactly %s characters", e.Param())
			case "numeric":
				errors[field] = "must contain only digits"
			case "uuid":
				errors[field] = "must be a valid uuid"
			default:
				errors[field] = "is invalid"
			}
		}

		logMsg := fmt.Sprint(errors)
		logger.Warn(logMsg)

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errors,
		})
		return
	}

	errorMsg := "Invalid request body"
	logger.Warn(errorMsg)

	c.JSON(http.StatusBadRequest, gin.H{
		"message": errorMsg,
	})
}
