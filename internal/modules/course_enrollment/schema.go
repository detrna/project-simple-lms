package course_enrollment

type CreateCourseEnrollmentRequest struct{}
type PatchCourseEnrollmentRequest struct {
	Score       *float64 `json:"score"`
	TeacherNote *string  `json:"teacherNote"`
}
