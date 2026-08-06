package seed

import (
	"context"
	"fmt"

	"main/internal/infrastructure/database"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func seedUsers(db *gorm.DB, ctx context.Context, state *State) error {
	password, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	adminPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	state.Users = make([]database.User, 0, 16)
	for i := 1; i <= 10; i++ {
		state.Users = append(state.Users, database.User{ID: uuid.MustParse(fmt.Sprintf("11111111-1111-1111-1111-%012d", 111111111110+i)), SystemID: fmt.Sprintf("STU-%03d", i), Name: fmt.Sprintf("Student %d", i), Email: fmt.Sprintf("student%d@example.com", i), Password: string(password), Role: "student"})
	}
	for i := 1; i <= 5; i++ {
		state.Users = append(state.Users, database.User{ID: uuid.MustParse(fmt.Sprintf("22222222-2222-2222-2222-%012d", 222222222220+i)), SystemID: fmt.Sprintf("INST-%03d", i), Name: fmt.Sprintf("Instructor %d", i), Email: fmt.Sprintf("instructor%d@example.com", i), Password: string(password), Role: "instructor"})
	}
	state.Users = append(state.Users, database.User{ID: uuid.MustParse("33333333-3333-3333-3333-333333333331"), SystemID: "ADMIN-001", Name: "Admin", Email: "admin@example.com", Password: string(adminPassword), Role: "admin"})
	return db.WithContext(ctx).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "system_id"}}, DoNothing: true}).Create(&state.Users).Error
}
