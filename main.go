package main

import (
	"flag"
	"fmt"
	"log"
	"main/config"
	"main/database"
	"main/internal/modules/template"
	"main/internal/modules/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	router := gin.Default()

	if err := database.Connect(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	if err := database.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	db := database.DB

	shouldSeed := flag.Bool("seed", false, "run database seeders")
	flag.Parse()
	if *shouldSeed {
		if err := database.Seed(db); err != nil {
			log.Fatal(err)
		}
		return
	}

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
	user.Register(api.Group("/users"), database.DB)

	router.Run()
}
