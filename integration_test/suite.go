package test_suite

import (
	"fmt"
	"log"

	"main/integration_test/factory"
	"main/internal/app"
	"main/internal/config"
	"main/internal/infrastructure"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupSuite() (*gin.Engine, *factory.Factory) {
	fmt.Print("TEST")
	cfg, err := config.Load()

	if err != nil {
		log.Fatal("couldn't load config")
	}

	infra, db, repo, err := infrastructure.Initialize(cfg)

	if err != nil {
		log.Fatal(err)
	}

	if err = TruncateDatabase(db); err != nil {
		log.Fatal("database error")
	}

	router := app.SetupRouter(cfg, infra, repo)
	factories := factory.NewFactory(infra, db, cfg)

	return router, factories
}

func TruncateDatabase(db *gorm.DB) error {
	return db.Exec(`
        TRUNCATE TABLE
            takes,
            classes,
            courses,
            users
        RESTART IDENTITY
        CASCADE;
    `).Error
}
