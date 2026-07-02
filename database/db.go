package database

import (
	"errors"
	"fmt"
	"main/internal/modules/instructor"
	"main/internal/modules/student"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
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

	return DB.AutoMigrate(&student.Student{}, &instructor.Instructor{})
}
