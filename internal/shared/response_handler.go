package shared

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseDTO[T any] struct {
	StatusCode *int
	Data       *T          `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

type ResponseSuccess[T any] struct {
	Data *T `json:"data"`
}

type PaginatedResponseSuccess[T any] struct {
	Data       *T          `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

func HandleResponse[T any](c *gin.Context, dto ResponseDTO[T]) {
	if dto.StatusCode == nil {
		statusOK := http.StatusOK
		dto.StatusCode = &statusOK
	}

	if dto.Pagination != nil {
		payload := PaginatedResponseSuccess[T]{
			Data:       dto.Data,
			Pagination: dto.Pagination,
		}

		c.JSON(*dto.StatusCode, payload)
		return
	} else {
		payload := ResponseSuccess[T]{
			Data: dto.Data,
		}

		c.JSON(*dto.StatusCode, payload)
		return
	}
}
