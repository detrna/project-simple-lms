package shared

import (
	"main/internal/pkg"

	"github.com/gin-gonic/gin"
)

func ParseJSON[T any](c *gin.Context, logger pkg.Logger) *T {
	var req T

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		HandleValidationError(c, logger, err)
		return nil
	}

	return &req
}

func ParseParams[T any](c *gin.Context, logger pkg.Logger) *T {
	var req T

	if err := c.ShouldBindUri(&req); err != nil {
		HandleValidationError(c, logger, err)
		return nil
	}

	return &req
}

func ParseQuery[T any](c *gin.Context, logger pkg.Logger) *T {
	var req T

	if err := c.ShouldBindQuery(&req); err != nil {
		HandleValidationError(c, logger, err)
		return nil
	}

	return &req
}
