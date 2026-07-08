package main

import (
	"flag"
	"log"
	"log/slog"

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

	logger := *infrastructure.Logger
	logger.Info("application starting")

	if *seed {
		database.Seed(infrastructure.DB)
	}

	port := infrastructure.Config.Server.Port

	router := app.SetupRouter(infrastructure)
	router.Run(":" + port)

	logger.Info("http server listening on port ", slog.String("port", port))
	defer logger.Info("application shutting down")
}
