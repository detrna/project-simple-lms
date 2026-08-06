package seed

import (
	"context"
	"time"

	"main/internal/infrastructure/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func seedCatalog(db *gorm.DB, ctx context.Context, state *State) error {
	courses := []database.Course{
		{ID: uuid.MustParse("44444444-4444-4444-4444-444444444441"), SystemID: "ENG-101", Name: "English", Credits: 3, AcademicYear: "2026/2027"},
		{ID: uuid.MustParse("44444444-4444-4444-4444-444444444442"), SystemID: "MAT-101", Name: "Mathematics", Credits: 3, AcademicYear: "2026/2027"},
		{ID: uuid.MustParse("44444444-4444-4444-4444-444444444443"), SystemID: "INF-101", Name: "Informatics", Credits: 3, AcademicYear: "2026/2027"},
		{ID: uuid.MustParse("44444444-4444-4444-4444-444444444444"), SystemID: "PE-101", Name: "Physical Education", Credits: 3, AcademicYear: "2026/2027"},
		{ID: uuid.MustParse("44444444-4444-4444-4444-444444444445"), SystemID: "PHY-101", Name: "Physics", Credits: 3, AcademicYear: "2026/2027"},
	}

	state.Classes = []database.Class{{ID: uuid.MustParse("55555555-5555-5555-5555-555555555551"), SystemID: "ENG-A", CourseID: courses[0].ID, Name: "English Class"}, {ID: uuid.MustParse("55555555-5555-5555-5555-555555555552"), SystemID: "MAT-A", CourseID: courses[1].ID, Name: "Mathematics Class"}, {ID: uuid.MustParse("55555555-5555-5555-5555-555555555553"), SystemID: "INF-A", CourseID: courses[2].ID, Name: "Informatics Class"}, {ID: uuid.MustParse("55555555-5555-5555-5555-555555555554"), SystemID: "PE-A", CourseID: courses[3].ID, Name: "Physical Education Class"}, {ID: uuid.MustParse("55555555-5555-5555-5555-555555555555"), SystemID: "PHY-A", CourseID: courses[4].ID, Name: "Physics Class"}}
	assignments := make([]database.Assignment, len(courses))
	for i, course := range courses {
		assignments[i] = database.Assignment{ID: uuid.MustParse("77777777-7777-7777-7777-77777777777" + string(rune('1'+i))), ClassID: state.Classes[i].ID, CourseID: course.ID, Title: course.Name + " Assignment", Description: "Complete the first course assignment.", Deadline: time.Now().AddDate(0, 0, 14)}
	}
	for _, rows := range []any{&courses, &state.Classes, &assignments} {
		if err := db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(rows).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedClassEnrollment(db *gorm.DB, ctx context.Context, state *State) error {
	enrollments, teaches := make([]database.CourseEnrollment, 5), make([]database.Teaches, 5)
	for i := 0; i < 5; i++ {
		enrollments[i] = database.CourseEnrollment{ID: uuid.MustParse("88888888-8888-8888-8888-88888888888" + string(rune('1'+i))), CourseID: state.Classes[i].CourseID, StudentID: state.Users[i].ID}
		teaches[i] = database.Teaches{ID: uuid.MustParse("66666666-6666-6666-6666-66666666666" + string(rune('1'+i))), UserID: state.Users[10+i].ID, ClassID: state.Classes[i].ID}
	}
	for _, rows := range []any{&enrollments, &teaches} {
		if err := db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(rows).Error; err != nil {
			return err
		}
	}
	return nil
}
