package seed

import (
	"context"

	"main/internal/infrastructure/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func seedLMS(db *gorm.DB, ctx context.Context, state *State) error {
	studentID, teacherID := state.Users[0].ID, state.Users[10].ID
	courseID, classID := state.Courses[0].ID, state.Classes[0].ID
	assignmentID := uuid.MustParse("77777777-7777-7777-7777-777777777771")
	announcementID := uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb1")
	transcriptID := uuid.MustParse("dddddddd-dddd-dddd-dddd-ddddddddddd1")
	semesterID := uuid.MustParse("ffffffff-ffff-ffff-ffff-fffffffffff1")
	enrollmentID := uuid.MustParse("eeeeeeee-eeee-eeee-eeee-eeeeeeeeeee1")
	teacherNote := "Enrolled for 2026/2027."
	materials := []database.Material{{ID: uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa1"), ClassID: classID, CourseID: courseID, Title: "Reading material", Description: "Week one reading material."}}
	announcements := []database.Announcement{{ID: announcementID, Title: "Welcome to English", Content: "Please review the reading material before class.", CourseID: courseID, TeacherID: teacherID}}
	submissions := []database.Submission{{ID: uuid.MustParse("cccccccc-cccc-cccc-cccc-ccccccccccc1"), AssignmentID: assignmentID, StudentID: studentID, StudentNote: "My completed answers.", TeacherNote: "Awaiting review."}}
	transcripts := []database.AcademicTranscript{{ID: transcriptID, StudentID: studentID}}
	semesters := []database.SemesterTranscript{{ID: semesterID, Period: "Semester 1", AcademicTranscriptID: transcriptID}}
	enrollments := []database.Enrollment{{ID: enrollmentID, CourseID: courseID, StudentID: studentID, ClassID: classID, Status: "active", AcademicYear: "2026/2027", TeacherNote: &teacherNote}}
	comments := []database.Comment{{ID: uuid.MustParse("13131313-1313-1313-1313-131313131313"), Content: "Looking forward to the class.", UserID: studentID, ParentType: "announcement", ParentID: announcementID}}
	files := []database.File{
		{ID: uuid.MustParse("99999999-9999-9999-9999-999999999991"), Name: "english-reading.pdf", FileURL: "https://example.com/files/english-reading.pdf", ContentType: "application/pdf", Size: 128000, Bucket: "material", ParentID: materials[0].ID},
		{ID: uuid.MustParse("99999999-9999-9999-9999-999999999992"), Name: "assignment-brief.pdf", FileURL: "https://example.com/files/assignment-brief.pdf", ContentType: "application/pdf", Size: 64000, Bucket: "assignment", ParentID: assignmentID},
		{ID: uuid.MustParse("99999999-9999-9999-9999-999999999993"), Name: "submission.pdf", FileURL: "https://example.com/files/submission.pdf", ContentType: "application/pdf", Size: 96000, Bucket: "submission", ParentID: submissions[0].ID},
		{ID: uuid.MustParse("99999999-9999-9999-9999-999999999994"), Name: "welcome.pdf", FileURL: "https://example.com/files/welcome.pdf", ContentType: "application/pdf", Size: 32000, Bucket: "announcement", ParentID: announcementID},
	}
	for _, rows := range []any{&materials, &announcements, &submissions, &transcripts, &semesters, &enrollments, &comments, &files} {
		if err := db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(rows).Error; err != nil {
			return err
		}
	}
	for _, join := range []struct {
		table  string
		values map[string]any
	}{{"course_teachers", map[string]any{"course_id": courseID, "user_id": teacherID}}, {"assignment_teachers", map[string]any{"assignment_id": assignmentID, "user_id": teacherID}}, {"semester_transcript_enrollments", map[string]any{"semester_transcript_id": semesterID, "enrollment_id": enrollmentID}}} {
		if err := db.WithContext(ctx).Table(join.table).Clauses(clause.OnConflict{DoNothing: true}).Create(join.values).Error; err != nil {
			return err
		}
	}
	return nil
}
