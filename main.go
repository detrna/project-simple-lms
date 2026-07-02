package main

import (
	"fmt"
	"main/config"
	"main/database"
	"main/internal/modules/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	router := gin.Default()

	database.Connect()
	database.Migrate()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/ping/:name", func(ctx *gin.Context) {
		var name string = ctx.Param("name")
		var message string = "Name: " + name

		ctx.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	router.GET("/welcome", func(ctx *gin.Context) {
		var firstname string = ctx.DefaultQuery("firstname", "guest")
		var lastname string = ctx.Query("lastname")

		var message string = fmt.Sprintf("hello %s %s", firstname, lastname)

		ctx.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	api := router.Group("/api/v1")

	template.Register(api.Group("/templates"), database.DB)

	router.Run()
}
