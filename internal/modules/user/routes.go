package user

import (
	"main/internal/middleware"
	"main/internal/pkg"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	controller    IController
	tokenProvider pkg.JWTProvider
	logger        pkg.Logger
}

func NewRoutes(c IController, tokenProvider pkg.JWTProvider, logger pkg.Logger) *Routes {
	return &Routes{controller: c, tokenProvider: tokenProvider, logger: logger}
}

func (routes Routes) RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/users")

	router.GET("/:id", routes.controller.GetUserByID)
	router.POST("", middleware.Authenticate(routes.tokenProvider), routes.controller.CreateUser)
	router.PATCH("/:id", middleware.Authenticate(routes.tokenProvider), routes.controller.AdminUpdateUser)
	router.DELETE("/:id", middleware.Authenticate(routes.tokenProvider), routes.controller.DeleteUser)
}
