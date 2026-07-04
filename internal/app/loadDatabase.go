package app

import (
	"flag"
	"log"
	"main/internal/database"
)

func LoadDatabase() {
	if err := database.Connect(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	if err := database.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	db := database.DB

	shouldSeed := flag.Bool("seed", false, "run database seeders")
	flag.Parse()

	if *shouldSeed {
		if err := database.Seed(db); err != nil {
			log.Fatal(err)
		}
		return
	}
}
