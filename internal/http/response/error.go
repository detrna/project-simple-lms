package response

import (
	"errors"
	"main/internal/pkg"
	apperrors "main/internal/shared/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func Error(c *gin.Context, logger pkg.Logger, err error) {
	logger.WarnSkip(1, err.Error())

	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		c.JSON(appErr.StatusCode, gin.H{
			"error":   err.Error(),
			"message": appErr.Message,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
