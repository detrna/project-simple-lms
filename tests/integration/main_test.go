package integration_test

import (
	"log"
	"os"
	"testing"

	"main/internal/app"
	"main/internal/config"
	"main/internal/infrastructure"
	"main/internal/infrastructure/database"
	"main/tests/factory"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine
var Factory *factory.Factory

func TestMain(m *testing.M) {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal("couldn't load config")
	}

	infra, db, repo, err := infrastructure.Initialize(*cfg)

	if err != nil {
		log.Fatal(err)
	}

	Factory = factory.NewFactory(infra, db, cfg)

	Router = app.SetupRouter(cfg, infra, repo)

	TruncateDatabase()
	defer TruncateDatabase()

	code := m.Run()

	os.Exit(code)
}

func TruncateDatabase() error {
	return database.DB.Exec(`
        TRUNCATE TABLE
            takes,
            classes,
            courses,
            users
        RESTART IDENTITY
        CASCADE;
    `).Error
}
