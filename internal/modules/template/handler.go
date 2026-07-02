package template

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase *UseCase
}

func NewHandler(usecase *UseCase) *Handler {
	return &Handler{usecase: usecase}
}

func (r *Handler) Ping(ctx *gin.Context) {
	result, err := r.usecase.repo.Ping()

	if err != nil {
		panic("Internal error")
	}

	ctx.JSON(http.StatusOK, gin.H{"message": result})
}
