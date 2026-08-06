package submission

import "github.com/google/uuid"

type CreateSubmissionRequest struct {
	TeacherNote string      `json:"teacherNote" binding:"required"`
	StudentNote string      `json:"studentNote" binding:"required"`
	FileIDs     []uuid.UUID `json:"fileIds" binding:"required"`
}

type StudentPatchSubmissionRequest struct {
	StudentNote *string      `json:"studentNote"`
	FileIDs     *[]uuid.UUID `json:"fileIds"`
}

type TeacherPatchSubmissionRequest struct {
	Score       *float64     `json:"score"`
	TeacherNote *string      `json:"teacherNote"`
	FileIDs     *[]uuid.UUID `json:"fileIds"`
}
