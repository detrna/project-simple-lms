package response

import (
	"main/internal/shared/pagination"

	"github.com/gin-gonic/gin"
)

type ResponseDTO[T any] struct {
	StatusCode int
	Data       T                              `json:"data"`
	Pagination *pagination.PaginationResponse `json:"pagination,omitempty"`
}

func Success[T any](c *gin.Context, dto ResponseDTO[T]) {
	if dto.Pagination != nil {
		c.JSON(dto.StatusCode, gin.H{
			"data":       dto.Data,
			"pagination": dto.Pagination,
		})
	} else {
		c.JSON(dto.StatusCode, gin.H{
			"data": dto.Data,
		})
	}
}
