package app

import (
	"main/internal/container"
	"main/internal/infrastructure"
	"main/internal/middleware"
	"main/internal/modules/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter(infra *infrastructure.Infrastructure) *gin.Engine {
	router := gin.Default()

	router.Use(
		middleware.RequestLogger(infra.Logger),
		middleware.ErrorLogger(infra.Logger),
	)

	userModule := container.NewUserContainer(infra)
	authModule := container.NewAuthContainer(infra, userModule.Repo)

	api := router.Group("/api/v1")
	user.RegisterRoutes(api, userModule.Controller)
	authModule.Routes.RegisterRoutes(api)

	return router
}
