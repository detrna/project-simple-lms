package template

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	usecase *UseCase
}

func NewController(usecase *UseCase) *Controller {
	return &Controller{usecase: usecase}
}

type IController interface {
	Ping() (string, error)
}

func (r *Controller) Ping(ctx *gin.Context) {
	result, err := r.usecase.repo.Ping()

	if err != nil {
		panic("Internal error")
	}

	ctx.JSON(http.StatusOK, gin.H{"message": result})
}
