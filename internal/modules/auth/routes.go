package auth

import (
	"main/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	controller IController
	secretKey  string
}

func NewRoutes(c IController, secretKey string) *Routes {
	return &Routes{controller: c, secretKey: secretKey}
}

func (routes Routes) RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.Use(middleware.Authenticate(routes.secretKey))

	router.POST("/login", routes.controller.Login)
	router.DELETE("/logout", routes.controller.Logout)
	router.POST("/refresh", routes.controller.Refresh)
}
