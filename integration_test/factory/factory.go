package factory

import (
	"main/internal/config"
	"main/internal/pkg"

	"gorm.io/gorm"
)

type Factory struct {
	Infra  *pkg.Packages
	DB     *gorm.DB
	Config *config.Config
}

func NewFactory(infra *pkg.Packages, db *gorm.DB, cfg *config.Config) *Factory {
	return &Factory{Infra: infra, DB: db, Config: cfg}
}
