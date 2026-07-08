package template

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, controller IController) {
	router.GET("", controller.Ping)
}
