package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, controller IController) {
	router := rg.Group("/users")

	router.GET("/:id", controller.GetUserByID)
	router.POST("", controller.CreateUser)
	router.PATCH("/:id", controller.UpdateUser)
	router.DELETE("/:id", controller.DeleteUser)
}
