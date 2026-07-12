package main

import (
	"flag"
	"log"
	"log/slog"

	"main/internal/app"
	"main/internal/config"
	"main/internal/infrastructure"
	"main/internal/infrastructure/database"
)

var seed = flag.Bool("seed", false, "run seed")

func main() {
	flag.Parse()

	cfg, err := config.Load()

	if err != nil {
		log.Fatal("couldn't load config")
	}

	packages, db, repository, err := infrastructure.Initialize(*cfg)

	if err != nil {
		log.Fatalf("Bootstrap failed: %v", err)
	}

	logger := packages.Logger
	logger.Info("application starting")

	if *seed {
		database.Seed(db)
	}

	port := cfg.Server.Port

	router := app.SetupRouter(*packages, *repository)
	router.Run(":" + port)

	logger.Info("http server listening on port ", slog.String("port", port))
	defer logger.Info("application shutting down")
}
