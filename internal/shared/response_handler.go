package shared

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseSuccess[T any] struct {
	Data       T           `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

func HandleResponse(c *gin.Context, payload ResponseSuccess[any]) {
	c.JSON(http.StatusOK, payload)
}
