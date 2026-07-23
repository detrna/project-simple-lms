package app

import (
	"main/internal/config"
	"main/internal/container"
	"main/internal/infrastructure/repository"
	"main/internal/middleware"
	"main/internal/pkg"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config, infra *pkg.Packages, repo *repository.Repository) *gin.Engine {
	router := gin.Default()

	router.Use(
		middleware.RequestLogger(infra.Logger),
		middleware.ErrorLogger(infra.Logger),
	)

	userModule := container.NewUserContainer(infra, repo)
	authModule := container.NewAuthContainer(cfg, infra, repo)

	api := router.Group("/api/v1")
	userModule.Routes.RegisterRoutes(api)
	authModule.Routes.RegisterRoutes(api)

	return router
}
