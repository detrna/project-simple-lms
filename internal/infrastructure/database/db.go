package database

import (
	"errors"
	"fmt"
	"log"
	"main/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dbConfig config.DatabaseConfig) error {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
		dbConfig.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func Migrate() error {
	if DB == nil {
		return errors.New("Database is not connected")
	}

	return DB.AutoMigrate(
		&User{},
		&Course{},
		&Takes{},
		&Teaches{},
		&Class{},
		&Material{},
		&MaterialFile{},
		&Assignment{},
		&AssignmentFile{},
		&SubmissionFile{},
		&SubmissionGrades{},
		&JWT{},
	)
}

func Load(dbConfig config.DatabaseConfig) *gorm.DB {
	if err := Connect(dbConfig); err != nil {
		log.Fatal(err)
	}

	if err := Migrate(); err != nil {
		log.Fatal(err)
	}

	return DB
}
