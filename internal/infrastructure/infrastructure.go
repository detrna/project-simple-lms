package infrastructure

import (
	"log/slog"
	"main/internal/config"
	"main/internal/infrastructure/database"

	"gorm.io/gorm"
)

type Infrastructure struct {
	Config *config.Config
	DB     *gorm.DB
	Logger *slog.Logger
}

func Initialize() (*Infrastructure, error) {
	cfg, err := config.Load()

	if err != nil {
		return nil, err
	}

	logger := NewLogger(cfg.Logger)

	db := database.Load(cfg.Database)

	logger.Info("database connected")

	return &Infrastructure{Config: cfg, DB: db, Logger: logger}, nil
}
