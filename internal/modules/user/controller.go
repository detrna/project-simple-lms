package user

import (
	"main/internal/shared"
	"net/http"

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
	GetUserByID(c *gin.Context)
	GetUserBySystemID(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func (controller *Controller) GetUserByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	ctx := c.Request.Context()

	result, err := controller.usecase.GetUserByID(ctx, id)

	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(200, result)
}

func (controller *Controller) GetUserBySystemID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("UserId"))

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	result, err := controller.usecase.GetUserByID(ctx, id)

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	user, err := controller.usecase.CreateUser(ctx, dto)

	if err != nil {
		shared.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (controller *Controller) UpdateUser(c *gin.Context) {
	var body UpdateUserBodySchema

	err := c.ShouldBindBodyWithJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	dto := UpdateUserSchema{
		ID:       &id,
		Name:     body.Name,
		Email:    body.Email,
		SystemID: body.SystemID,
	}

	ctx := c.Request.Context()

	result, err := controller.usecase.UpdateUser(ctx, dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

func (controller *Controller) DeleteUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()

	if err := controller.usecase.DeleteUser(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusNoContent, "")
}
