package database

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var Users []User
var Classes []Class

func Seed(db *gorm.DB) error {
	if db == nil {
		return errors.New("Database is not connected")
	}

	ctx := context.Background()

	if err := seedUsers(db, ctx); err != nil {
		return err
	}

	if err := seedCourses(db, ctx); err != nil {
		return err
	}

	if err := seedClasses(db, ctx); err != nil {
		return err
	}

	if err := seedTakes(db, ctx); err != nil {
		return err
	}

	if err := seedAssignments(db, ctx); err != nil {
		return err
	}

	if err := seedTeaches(db, ctx); err != nil {
		return err
	}

	return nil
}

func seedUsers(db *gorm.DB, ctx context.Context) error {
	Users = []User{
		{ID: uuid.MustParse("11111111-1111-1111-1111-111111111111"), SystemID: "STU-001", Name: "Student 1", Email: "student1@example.com", Password: "password", Role: "student"},
		{ID: uuid.MustParse("11111111-1111-1111-1111-111111111112"), SystemID: "STU-002", Name: "Student 2", Email: "student2@example.com", Password: "password", Role: "student"},
		{ID: uuid.MustParse("11111111-1111-1111-1111-111111111113"), SystemID: "STU-003", Name: "Student 3", Email: "student3@example.com", Password: "password", Role: "student"},
		{ID: uuid.MustParse("11111111-1111-1111-1111-111111111114"), SystemID: "STU-004", Name: "Student 4", Email: "student4@example.com", Password: "password", Role: "student"},
		{ID: uuid.MustParse("11111111-1111-1111-1111-111111111115"), SystemID: "STU-005", Name: "Student 5", Email: "student5@example.com", Password: "password", Role: "student"},
		{ID: uuid.MustParse("11111111-1111-1111-1111-111111111116"), SystemID: "STU-006", Name: "Student 6", Email: "student6@example.com", Password: "password", Role: "student"},
		{ID: uuid.MustParse("11111111-1111-1111-1111-111111111117"), SystemID: "STU-007", Name: "Student 7", Email: "student7@example.com", Password: "password", Role: "student"},
		{ID: uuid.MustParse("11111111-1111-1111-1111-111111111118"), SystemID: "STU-008", Name: "Student 8", Email: "student8@example.com", Password: "password", Role: "student"},
		{ID: uuid.MustParse("11111111-1111-1111-1111-111111111119"), SystemID: "STU-009", Name: "Student 9", Email: "student9@example.com", Password: "password", Role: "student"},
		{ID: uuid.MustParse("11111111-1111-1111-1111-111111111120"), SystemID: "STU-010", Name: "Student 10", Email: "student10@example.com", Password: "password", Role: "student"},
		{ID: uuid.MustParse("22222222-2222-2222-2222-222222222221"), SystemID: "INST-001", Name: "Instructor 1", Email: "instructor1@example.com", Password: "password", Role: "instructor"},
		{ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"), SystemID: "INST-002", Name: "Instructor 2", Email: "instructor2@example.com", Password: "password", Role: "instructor"},
		{ID: uuid.MustParse("22222222-2222-2222-2222-222222222223"), SystemID: "INST-003", Name: "Instructor 3", Email: "instructor3@example.com", Password: "password", Role: "instructor"},
		{ID: uuid.MustParse("22222222-2222-2222-2222-222222222224"), SystemID: "INST-004", Name: "Instructor 4", Email: "instructor4@example.com", Password: "password", Role: "instructor"},
		{ID: uuid.MustParse("22222222-2222-2222-2222-222222222225"), SystemID: "INST-005", Name: "Instructor 5", Email: "instructor5@example.com", Password: "password", Role: "instructor"},
		{ID: uuid.MustParse("33333333-3333-3333-3333-333333333331"), SystemID: "ADMIN-001", Name: "Admin", Email: "admin@example.com", Password: "password", Role: "admin"},
	}

	return db.WithContext(ctx).
		Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "system_id"}}, DoNothing: true}).
		Create(&Users).Error
}

func seedCourses(db *gorm.DB, ctx context.Context) error {
	courses := []Course{
		{ID: uuid.MustParse("44444444-4444-4444-4444-444444444441"), Name: "English", Credits: 3},
		{ID: uuid.MustParse("44444444-4444-4444-4444-444444444442"), Name: "Mathematics", Credits: 3},
		{ID: uuid.MustParse("44444444-4444-4444-4444-444444444443"), Name: "Informatics", Credits: 3},
		{ID: uuid.MustParse("44444444-4444-4444-4444-444444444444"), Name: "Physical Education", Credits: 3},
		{ID: uuid.MustParse("44444444-4444-4444-4444-444444444445"), Name: "Physics", Credits: 3},
	}

	return db.WithContext(ctx).
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(&courses).Error
}

func seedClasses(db *gorm.DB, ctx context.Context) error {
	Classes = []Class{
		{ID: uuid.MustParse("55555555-5555-5555-5555-555555555551"), CourseID: uuid.MustParse("44444444-4444-4444-4444-444444444441"), Name: "English Class"},
		{ID: uuid.MustParse("55555555-5555-5555-5555-555555555552"), CourseID: uuid.MustParse("44444444-4444-4444-4444-444444444442"), Name: "Mathematics Class"},
		{ID: uuid.MustParse("55555555-5555-5555-5555-555555555553"), CourseID: uuid.MustParse("44444444-4444-4444-4444-444444444443"), Name: "Informatics Class"},
		{ID: uuid.MustParse("55555555-5555-5555-5555-555555555554"), CourseID: uuid.MustParse("44444444-4444-4444-4444-444444444444"), Name: "Physical Education Class"},
		{ID: uuid.MustParse("55555555-5555-5555-5555-555555555555"), CourseID: uuid.MustParse("44444444-4444-4444-4444-444444444445"), Name: "Physics Class"},
	}

	return db.WithContext(ctx).
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(&Classes).Error
}

func seedTakes(db *gorm.DB, ctx context.Context) error {
	takes := []Takes{
		{ID: uuid.MustParse("88888888-8888-8888-8888-888888888881"), UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"), ClassID: uuid.MustParse("55555555-5555-5555-5555-555555555551"), Grade: 0},
		{ID: uuid.MustParse("88888888-8888-8888-8888-888888888882"), UserID: uuid.MustParse("11111111-1111-1111-1111-111111111112"), ClassID: uuid.MustParse("55555555-5555-5555-5555-555555555551"), Grade: 0},
		{ID: uuid.MustParse("88888888-8888-8888-8888-888888888883"), UserID: uuid.MustParse("11111111-1111-1111-1111-111111111113"), ClassID: uuid.MustParse("55555555-5555-5555-5555-555555555552"), Grade: 0},
		{ID: uuid.MustParse("88888888-8888-8888-8888-888888888884"), UserID: uuid.MustParse("11111111-1111-1111-1111-111111111114"), ClassID: uuid.MustParse("55555555-5555-5555-5555-555555555553"), Grade: 0},
		{ID: uuid.MustParse("88888888-8888-8888-8888-888888888885"), UserID: uuid.MustParse("11111111-1111-1111-1111-111111111115"), ClassID: uuid.MustParse("55555555-5555-5555-5555-555555555554"), Grade: 0},
	}

	return db.WithContext(ctx).
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(&takes).Error
}

func seedAssignments(db *gorm.DB, ctx context.Context) error {
	assignments := []Assignment{
		{ID: uuid.MustParse("77777777-7777-7777-7777-777777777771"), ClassID: uuid.MustParse("55555555-5555-5555-5555-555555555551"), Title: "English Assignment", Description: "Read the short story and answer questions.", Deadline: time.Now().AddDate(0, 0, 14)},
		{ID: uuid.MustParse("77777777-7777-7777-7777-777777777772"), ClassID: uuid.MustParse("55555555-5555-5555-5555-555555555552"), Title: "Mathematics Assignment", Description: "Solve the algebra worksheet.", Deadline: time.Now().AddDate(0, 0, 14)},
		{ID: uuid.MustParse("77777777-7777-7777-7777-777777777773"), ClassID: uuid.MustParse("55555555-5555-5555-5555-555555555553"), Title: "Informatics Assignment", Description: "Build a basic data structure.", Deadline: time.Now().AddDate(0, 0, 14)},
		{ID: uuid.MustParse("77777777-7777-7777-7777-777777777774"), ClassID: uuid.MustParse("55555555-5555-5555-5555-555555555554"), Title: "Physical Education Assignment", Description: "Prepare a one-week physical activity plan.", Deadline: time.Now().AddDate(0, 0, 14)},
		{ID: uuid.MustParse("77777777-7777-7777-7777-777777777775"), ClassID: uuid.MustParse("55555555-5555-5555-5555-555555555555"), Title: "Physics Assignment", Description: "Complete the motion experiment report.", Deadline: time.Now().AddDate(0, 0, 14)},
	}

	return db.WithContext(ctx).
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(&assignments).Error
}

func seedTeaches(db *gorm.DB, ctx context.Context) error {
	teaches := []Teaches{
		{ID: uuid.MustParse("66666666-6666-6666-6666-666666666661"), UserID: uuid.MustParse("22222222-2222-2222-2222-222222222221"), ClassID: uuid.MustParse("55555555-5555-5555-5555-555555555551")},
		{ID: uuid.MustParse("66666666-6666-6666-6666-666666666662"), UserID: uuid.MustParse("22222222-2222-2222-2222-222222222222"), ClassID: uuid.MustParse("55555555-5555-5555-5555-555555555552")},
		{ID: uuid.MustParse("66666666-6666-6666-6666-666666666663"), UserID: uuid.MustParse("22222222-2222-2222-2222-222222222223"), ClassID: uuid.MustParse("55555555-5555-5555-5555-555555555553")},
		{ID: uuid.MustParse("66666666-6666-6666-6666-666666666664"), UserID: uuid.MustParse("22222222-2222-2222-2222-222222222224"), ClassID: uuid.MustParse("55555555-5555-5555-5555-555555555554")},
		{ID: uuid.MustParse("66666666-6666-6666-6666-666666666665"), UserID: uuid.MustParse("22222222-2222-2222-2222-222222222225"), ClassID: uuid.MustParse("55555555-5555-5555-5555-555555555555")},
	}

	return db.WithContext(ctx).
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(&teaches).Error
}
