package user

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
	GetUserById(c *gin.Context)
	GetUserByUserId(c *gin.Context)
	CreateUser(c *gin.Context)
}

func (controller *Controller) GetUserById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	ctx := c.Request.Context()

	result, err := controller.usecase.GetUserById(ctx, id)

	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(200, result)
}

func (controller *Controller) GetUserByUserId(c *gin.Context) {
	id, err := uuid.Parse(c.Param("UserId"))

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	result, err := controller.usecase.GetUserById(ctx, id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

func (controller *Controller) CreateUser(c *gin.Context) {
	var dto CreateUserSchema

	err := c.ShouldBindBodyWithJSON(&dto)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	err = controller.usecase.CreateUser(ctx, dto)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully"})
}
