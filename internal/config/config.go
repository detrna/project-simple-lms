package config

import "github.com/joho/godotenv"

func Load() {
	if err := godotenv.Load(); err == nil {
		return
	}

	if err := godotenv.Load("../../.env"); err == nil {
		return
	}

	if err := godotenv.Load("../../../.env"); err == nil {
		return
	}

	panic("could not load .env")
}
