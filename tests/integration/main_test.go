package integration_test

import (
	"log"
	"os"
	"testing"

	"main/internal/app"
	"main/internal/infrastructure"
	"main/internal/infrastructure/database"
	"main/tests/factory"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func TestMain(m *testing.M) {
	infra, err := infrastructure.Initialize()

	if err != nil {
		log.Fatal(err)
	}

	factory.Infra = infra
	factory.DB = infra.DB

	Router = app.SetupRouter(infra)

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
