package app

import (
	"main/internal/container"
	"main/internal/infrastructure"
	"main/internal/modules/auth"
	"main/internal/modules/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter(infra *infrastructure.Infrastructure) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")

	userModule := container.NewUserContainer(infra)
	authModule := container.NewAuthContainer(infra, userModule.Repo)

	user.RegisterRoutes(api, userModule.Controller)
	auth.RegisterRoutes(api, authModule.Controller)

	return router
}
