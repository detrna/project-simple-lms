package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, controller *Controller) {
	router.GET("/:id", controller.GetUserByID)
	router.POST("", controller.CreateUser)
	router.PATCH("/:id", controller.UpdateUser)
	router.DELETE("/:id", controller.DeleteUser)
}
