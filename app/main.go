package main

import (
	"flag"
	"log"

	"main/internal/app"
	"main/internal/infrastructure"
	"main/internal/infrastructure/database"
)

var seed = flag.Bool("seed", false, "run seed")

func main() {
	flag.Parse()

	infrastructure, err := infrastructure.Initialize()

	if err != nil {
		log.Fatalf("Bootstrap failed: %v", err)
	}

	if *seed {
		database.Seed(infrastructure.DB)
	}

	port := infrastructure.Config.Server.Port
	router := app.SetupRouter(infrastructure)
	router.Run(port)
}
