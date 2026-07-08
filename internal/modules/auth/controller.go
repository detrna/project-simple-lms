package auth

import (
	"main/internal/shared"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	usecase *UseCase
	logger  shared.Logger
}

func NewController(usecase *UseCase, logger shared.Logger) *Controller {
	return &Controller{usecase: usecase, logger: logger}
}

type IController interface {
	Login(c *gin.Context)
	// Recover(c *gin.Context)
	// Logout(c *gin.Context)
	// Refresh(c *gin.Context)
}

func (controller *Controller) Login(c *gin.Context) {
	var dto LoginSchema

	err := c.ShouldBindBodyWithJSON(&dto)

	if err != nil {
		shared.HandleError(c, controller.logger, shared.ErrBadRequest)
		return
	}

	ctx := c.Request.Context()

	Tokens, err := controller.usecase.Login(ctx, dto)

	if err != nil {
		shared.HandleError(c, controller.logger, err)
		return
	}

	CookieEnv := os.Getenv("COOKIE_SECURE")
	var secure bool

	if CookieEnv == "true" {
		secure = true
	} else {
		secure = false
	}

	c.SetCookie(
		"access_token",      // name
		Tokens.RefreshToken, // value
		3600,                // maxAge (seconds)
		"/",                 // path
		"",                  // domain
		secure,              // secure
		true,                // httpOnly
	)

	c.JSON(http.StatusOK, Tokens.AccessToken)
}
