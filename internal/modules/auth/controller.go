package auth

import (
	"main/internal/config"
	"main/internal/pkg"
	"main/internal/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	usecase *UseCase
	logger  pkg.Logger
	cfg     *config.Config
}

func NewController(usecase *UseCase, logger pkg.Logger) *Controller {
	return &Controller{usecase: usecase, logger: logger}
}

type IController interface {
	Login(c *gin.Context)
	// Recover(c *gin.Context)
	Logout(c *gin.Context)
	Refresh(c *gin.Context)
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

	secure := controller.cfg.App.Mode == "PRODUCTION"

	c.SetCookie(
		"refresh_token",     // name
		Tokens.RefreshToken, // value
		3600,                // maxAge (seconds)
		"/",                 // path
		"",                  // domain
		secure,              // secure
		true,                // httpOnly
	)

	c.JSON(http.StatusOK, Tokens.AccessToken)
}

func (controller *Controller) Logout(c *gin.Context) {
	value, _ := c.Get("user")
	JWTPayload, _ := value.(*JWT)

	ctx := c.Request.Context()
	err := controller.usecase.Logout(ctx, JWTPayload.JTI)

	if err != nil {
		shared.HandleError(c, controller.logger, err)
	}

	secure := controller.cfg.App.Mode == "PRODUCTION"

	c.SetCookie(
		"refresh_token", // name
		"",              // value
		-1,              // maxAge (seconds)
		"/",             // path
		"",              // domain
		secure,          // secure
		true,            // httpOnly
	)

	c.JSON(http.StatusOK, gin.H{"message": "successfully logged out"})
}

func (controller *Controller) Refresh(c *gin.Context) {
	httpToken, err := c.Request.Cookie("refresh_token")
	token := httpToken.Value

	if err != nil {
		c.JSON(http.StatusUnauthorized, "token did not exist")
	}

	key := controller.cfg.JWT.RefreshSecret
	claims, err := shared.ParseToken(token, key)

	if err != nil {
		c.JSON(http.StatusUnauthorized, "token expired")
	}

	ctx := c.Request.Context()
	result, err := controller.usecase.Refresh(ctx, claims.Payload)

	secure := controller.cfg.App.Mode == "PRODUCTION"
	c.SetCookie(
		"access_token",     // name
		result.AccessToken, // value
		3600,               // maxAge (seconds)
		"/",                // path
		"",                 // domain
		secure,             // secure
		true,               // httpOnly
	)

	c.JSON(http.StatusOK, result.AccessToken)
}
