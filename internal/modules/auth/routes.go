package auth

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

func NewRoutes(c IController, jwtProvider pkg.JWTProvider, logger pkg.Logger) *Routes {
	return &Routes{controller: c, tokenProvider: jwtProvider, logger: logger}
}

func (routes Routes) RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/login", routes.controller.Login)
	router.DELETE("/logout", middleware.Authenticate(routes.tokenProvider), routes.controller.Logout)
	router.POST("/refresh", routes.controller.Refresh)
}
