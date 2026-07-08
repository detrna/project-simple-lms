package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, controller IController) {
	router := rg.Group("/auth")

	router.GET("/login", controller.Login)
}
