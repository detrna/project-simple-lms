package auth

import (
	"main/internal/http/validator"
	"main/internal/pkg"
	"main/internal/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	isProduction bool
	usecase      IUseCase
	logger       pkg.Logger
}

func NewController(usecase IUseCase, logger pkg.Logger, isProduction bool) *Controller {
	return &Controller{usecase: usecase, logger: logger, isProduction: isProduction}
}

type IController interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Refresh(c *gin.Context)
	Recover(c *gin.Context)
	VerifyRecovery(c *gin.Context)
}

func (controller *Controller) Login(c *gin.Context) {
	body := validator.ValidateBody[LoginSchema](c, controller.logger)

	ctx := c.Request.Context()
	Tokens, err := controller.usecase.Login(ctx, body)

	if err != nil {
		shared.HandleError(c, controller.logger, err)
		return
	}

	c.SetCookie(
		"refresh_token",         // name
		Tokens.RefreshToken,     // value
		3600,                    // maxAge (seconds)
		"/",                     // path
		"",                      // domain
		controller.isProduction, // secure
		true,                    // httpOnly
	)

	payload := shared.ResponseDTO[TokenResponse]{
		Data: &TokenResponse{
			AccessToken: Tokens.AccessToken,
		},
	}

	shared.HandleResponse(c, payload)
}

func (controller *Controller) Logout(c *gin.Context) {
	token, err := c.Cookie("refresh_token")

	if err != nil {
		shared.HandleError(c, controller.logger, shared.ErrUnauthorized)
		return
	}

	ctx := c.Request.Context()
	err = controller.usecase.Logout(ctx, token)

	if err != nil {
		shared.HandleError(c, controller.logger, err)
	}

	c.SetCookie(
		"refresh_token",         // name
		"",                      // value
		-1,                      // maxAge (seconds)
		"/",                     // path
		"",                      // domain
		controller.isProduction, // secure
		true,                    // httpOnly
	)

	c.JSON(http.StatusNoContent, nil)
}

func (controller *Controller) Refresh(c *gin.Context) {
	token, err := c.Cookie("refresh_token")

	if err != nil {
		c.JSON(http.StatusUnauthorized, "token did not exist")
		return
	}

	ctx := c.Request.Context()
	result, err := controller.usecase.Refresh(ctx, token)

	if err != nil {
		shared.HandleError(c, controller.logger, err)
		return
	}

	c.SetCookie(
		"refresh_token",         // name
		result.RefreshToken,     // value
		3600,                    // maxAge (seconds)
		"/",                     // path
		"",                      // domain
		controller.isProduction, // secure
		true,                    // httpOnly
	)

	payload := shared.ResponseDTO[TokenResponse]{
		Data: &TokenResponse{AccessToken: result.AccessToken},
	}

	shared.HandleResponse(c, payload)
}

func (controller *Controller) Recover(c *gin.Context) {
	body := validator.ValidateBody[RecoverSchema](c, controller.logger)

	ctx := c.Request.Context()
	err := controller.usecase.Recover(ctx, body)

	if err != nil {
		shared.HandleError(c, controller.logger, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (controller *Controller) VerifyRecovery(c *gin.Context) {
	body := validator.ValidateBody[VerifyRecoverySchema](c, controller.logger)

	ctx := c.Request.Context()
	err := controller.usecase.VerifyRecovery(ctx, body)

	if err != nil {
		shared.HandleError(c, controller.logger, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
