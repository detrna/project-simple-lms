package shared

import (
	"main/internal/shared/pagination"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseDTO[T any] struct {
	StatusCode *int
	Data       *T                             `json:"data"`
	Pagination *pagination.PaginationResponse `json:"pagination"`
}

type ResponseSuccess[T any] struct {
	Data *T `json:"data"`
}

type PaginatedResponseSuccess[T any] struct {
	Data       *[]T                           `json:"data"`
	Pagination *pagination.PaginationResponse `json:"pagination"`
}

func HandleResponse[T any](c *gin.Context, dto ResponseDTO[T]) {
	if dto.StatusCode == nil {
		statusOK := http.StatusOK
		dto.StatusCode = &statusOK
	}
}
