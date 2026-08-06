package seed

import (
	"context"
	"errors"

	"main/internal/infrastructure/database"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	if db == nil {
		return errors.New("database is not connected")
	}
	ctx := context.Background()
	state := &State{}
	for _, seed := range []func(*gorm.DB, context.Context, *State) error{
		seedUsers,
		seedCatalog,
		seedLMS,
		seedClassEnrollment,
	} {
		if err := seed(db, ctx, state); err != nil {
			return err
		}
	}
	return nil
}

type State struct {
	Users   []database.User
	Classes []database.Class
}
