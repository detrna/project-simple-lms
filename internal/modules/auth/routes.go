package auth

import (
	"main/internal/middleware"
	"main/internal/pkg"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	controller  IController
	jwtProvider pkg.JWTProvider
}

func NewRoutes(c IController, jwtProvider pkg.JWTProvider) *Routes {
	return &Routes{controller: c, jwtProvider: jwtProvider}
}

func (routes Routes) RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.Use(middleware.Authenticate(routes.jwtProvider))

	router.POST("/login", routes.controller.Login)
	router.DELETE("/logout", routes.controller.Logout)
	router.POST("/refresh", routes.controller.Refresh)
}
