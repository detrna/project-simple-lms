package app

import (
	"main/internal/database"
	"main/internal/modules/class"
	"main/internal/modules/template"
	"main/internal/modules/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")

	template.Register(api.Group("/templates"), database.DB)
	user.Register(api.Group("/users"), database.DB)
	class.Register(api.Group("/classes"), database.DB)

	return router
}
