package class

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller struct {
	usecase *UseCase
}

func NewController(usecase *UseCase) *Controller {
	return &Controller{usecase: usecase}
}

type IController interface {
	GetStudents(c *gin.Context)
	GetMyClasses(c *gin.Context)
}

func (controller Controller) GetStudents(c *gin.Context) {
	classID, err := uuid.Parse(c.Param("classId"))

	if err != nil {
		c.JSON(404, gin.H{"Error": err.Error()})
	}

	ctx := c.Request.Context()

	result, err := controller.usecase.GetStudents(ctx, classID)

	c.JSON(200, result)
}

func (controller Controller) GetMyClasses(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("userId"))

	if err != nil {
		c.JSON(404, gin.H{"Error": err.Error()})
	}

	ctx := c.Request.Context()

	result, err := controller.usecase.GetMyClasses(ctx, userID)

	c.JSON(200, result)
}
