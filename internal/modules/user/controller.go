package user

import (
	"fmt"
	"main/internal/pkg"
	"main/internal/shared"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller struct {
	usecase IUseCase
	logger  pkg.Logger
}

func NewController(usecase IUseCase, logger pkg.Logger) *Controller {
	return &Controller{usecase: usecase, logger: logger}
}

type IController interface {
	GetUserByID(c *gin.Context)
	GetUserBySystemID(c *gin.Context)
	GetMyAccount(c *gin.Context)
	CreateUser(c *gin.Context)
	AdminUpdateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func (controller *Controller) GetUserByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		shared.HandleError(c, controller.logger, shared.ErrBadRequest)
		return
	}

	ctx := c.Request.Context()
	result, err := controller.usecase.GetUserByID(ctx, id)

	if err != nil {
		shared.HandleError(c, controller.logger, err)
		return
	}

	payload := shared.ResponseDTO[UserResponse]{
		Data: result,
	}

	shared.HandleResponse(c, payload)
}

func (controller *Controller) GetUserBySystemID(c *gin.Context) {
	id := c.Param("systemId")

	if id == "" {
		shared.HandleError(c, controller.logger, shared.ErrBadRequest)
		return
	}

	ctx := c.Request.Context()
	result, err := controller.usecase.GetUserBySystemID(ctx, id)

	if err != nil {
		fmt.Print("OALAH")
		shared.HandleError(c, controller.logger, err)
		return
	}

	payload := shared.ResponseDTO[UserResponse]{
		Data: result,
	}

	shared.HandleResponse(c, payload)
}

func (controller *Controller) GetMyAccount(c *gin.Context) {
	user, err := shared.GetJWTPayload(c)

	if err != nil {
		shared.HandleError(c, controller.logger, err)
		return
	}

	ctx := c.Request.Context()
	result, err := controller.usecase.GetUserByID(ctx, user.UserID)

	if err != nil {
		shared.HandleError(c, controller.logger, err)
		return
	}

	payload := shared.ResponseDTO[UserResponse]{
		Data: result,
	}

	shared.HandleResponse(c, payload)
}

func (controller *Controller) CreateUser(c *gin.Context) {
	var body CreateUserSchema
	err := c.ShouldBindBodyWithJSON(&body)

	if err != nil {
		shared.HandleError(c, controller.logger, err)
		return
	}

	ctx := c.Request.Context()
	result, err := controller.usecase.CreateUser(ctx, &body)

	if err != nil {
		shared.HandleError(c, controller.logger, err)
		return
	}

	statusCode := http.StatusCreated
	payload := shared.ResponseDTO[UserResponse]{
		Data:       result,
		StatusCode: &statusCode,
	}

	shared.HandleResponse(c, payload)
}

func (controller *Controller) AdminUpdateUser(c *gin.Context) {
	var body AdminUpdateUserSchema
	err := c.ShouldBindBodyWithJSON(&body)

	if err != nil {
		shared.HandleError(c, controller.logger, err)
		return
	}

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		shared.HandleError(c, controller.logger, err)
		return
	}

	dto := AdminUpdateUserDTO{
		ID:       id,
		Name:     body.Name,
		Email:    body.Email,
		SystemID: body.SystemID,
	}

	ctx := c.Request.Context()
	result, err := controller.usecase.AdminUpdateUser(ctx, &dto)

	if err != nil {
		shared.HandleError(c, controller.logger, err)
		return
	}

	payload := shared.ResponseDTO[UserResponse]{
		Data: result,
	}

	shared.HandleResponse(c, payload)
}

func (controller *Controller) DeleteUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		shared.HandleError(c, controller.logger, err)
		return
	}

	ctx := c.Request.Context()
	if err := controller.usecase.DeleteUser(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusNoContent, "")
}

func (controller Controller) UpdateUser(c *gin.Context) {
	user, err := shared.GetJWTPayload(c)

	if err != nil {
		shared.HandleError(c, controller.logger, nil)
		return
	}

	var body UpdateUserSchema
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		shared.HandleError(c, controller.logger, err)
		return
	}

	dto := UpdateUserDTO{
		User:     user,
		Password: body.Password,
	}

	ctx := c.Request.Context()
	result, err := controller.usecase.UpdateUser(ctx, &dto)

	if err != nil {
		shared.HandleError(c, controller.logger, err)
		return
	}

	payload := shared.ResponseDTO[UserResponse]{
		Data: result,
	}

	shared.HandleResponse(c, payload)
}
