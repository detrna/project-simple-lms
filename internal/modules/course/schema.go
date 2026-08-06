package course

import "github.com/google/uuid"

type CreateCourseRequest struct {
	SystemID     string      `json:"systemId" binding:"required"`
	Name         string      `json:"name" binding:"required"`
	Credits      int         `json:"credits" binding:"required"`
	AcademicYear string      `json:"academicYear" binding:"required"`
	TeacherIDs   []uuid.UUID `json:"teacherIds" binding:"required"`
}
type PatchCourseRequest struct {
	SystemID     *string      `json:"systemId"`
	Name         *string      `json:"name"`
	Credits      *int         `json:"credits"`
	AcademicYear *string      `json:"academicYear"`
	TeacherIDs   *[]uuid.UUID `json:"teacherIds"`
}
