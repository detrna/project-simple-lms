package main

import (
	"log"
	"main/internal/app"
)

func main() {
	if err := app.Initialize(); err != nil {
		log.Fatal(err)
	}

	router := app.NewApp()

	router.Run(":8080")
}
