package class

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	usecase *UseCase
}

func NewController(usecase *UseCase) *Controller {
	return &Controller{usecase: usecase}
}

type IController interface {
	GetClasses(c *gin.Context)
	GetClassByID(c *gin.Context)
	GetClassBySystemID(c *gin.Context)
	CreateClass(c *gin.Context)
	UpdateClass(c *gin.Context)
	DeleteClass(c *gin.Context)
}

func (controller Controller) GetClasses(c *gin.Context) {

}

func (controller Controller) GetClassByID(c *gin.Context) {

}

func (controller Controller) GetClassBySystemID(c *gin.Context) {

}

func (controller Controller) CreateClass(c *gin.Context) {

}

func (controller Controller) UpdateClass(c *gin.Context) {

}

func (controller Controller) DeleteClass(c *gin.Context) {

}
