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

	state.Courses = courses
	state.Courses = courses
	state.Classes = []database.Class{{ID: uuid.MustParse("55555555-5555-5555-5555-555555555551"), SystemID: "ENG-A", Name: "English Class"}, {ID: uuid.MustParse("55555555-5555-5555-5555-555555555552"), SystemID: "MAT-A", Name: "Mathematics Class"}, {ID: uuid.MustParse("55555555-5555-5555-5555-555555555553"), SystemID: "INF-A", Name: "Informatics Class"}, {ID: uuid.MustParse("55555555-5555-5555-5555-555555555554"), SystemID: "PE-A", Name: "Physical Education Class"}, {ID: uuid.MustParse("55555555-5555-5555-5555-555555555555"), SystemID: "PHY-A", Name: "Physics Class"}}
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
	enrollments := make([]database.Enrollment, 0, len(state.Classes))
	for i := 0; i < len(state.Classes) && i < len(state.Courses) && i < len(state.Users); i++ {
		enrollments = append(enrollments, database.Enrollment{
			ID:           uuid.MustParse("88888888-8888-8888-8888-88888888888" + string(rune('1'+i))),
			CourseID:     state.Courses[i].ID,
			StudentID:    state.Users[i].ID,
			ClassID:      state.Classes[i].ID,
			Status:       "active",
			AcademicYear: "2026/2027",
		})
	}
	if err := db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&enrollments).Error; err != nil {
		return err
	}
	return nil
}
