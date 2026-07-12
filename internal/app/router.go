package app

import (
	"main/internal/container"
	"main/internal/infrastructure/repository"
	"main/internal/middleware"
	"main/internal/modules/user"
	"main/internal/pkg"

	"github.com/gin-gonic/gin"
)

func SetupRouter(infra pkg.Packages, repo repository.Repository) *gin.Engine {
	router := gin.Default()

	router.Use(
		middleware.RequestLogger(infra.Logger),
		middleware.ErrorLogger(infra.Logger),
	)

	userModule := container.NewUserContainer(infra, repo)
	authModule := container.NewAuthContainer(infra, repo)

	api := router.Group("/api/v1")
	user.RegisterRoutes(api, userModule.Controller)
	authModule.Routes.RegisterRoutes(api)

	return router
}
