package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, controller *Controller) {
	router.GET("/:id", controller.GetUserById)
	router.GET("/:userId", controller.GetUserByUserId)
	router.POST("", controller.CreateUser)
}
