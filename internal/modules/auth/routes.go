package auth

import (
	"main/internal/config"
	"main/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	controller IController
	cfg        config.Config
}

func NewRoutes(c IController, cfg config.Config) *Routes {
	return &Routes{controller: c, cfg: cfg}
}

func (routes Routes) RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.Use(middleware.Authenticate(routes.cfg.JWT))

	router.POST("/login", routes.controller.Login)
	router.DELETE("/logout", routes.controller.Logout)
	router.POST("/refresh", routes.controller.Refresh)
}
