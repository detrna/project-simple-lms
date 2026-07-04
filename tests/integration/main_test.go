package integration_test

import (
	"log"
	"os"
	"testing"

	"main/internal/app"
	"main/internal/database"
)

func TestMain(m *testing.M) {
	if err := app.Initialize(); err != nil {
		log.Fatal(err)
	}

	TruncateDatabase()

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
