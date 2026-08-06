package assignment

import (
	"github.com/google/uuid"
	"time"
)

type CreateAssignmentRequest struct {
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Deadline    time.Time   `json:"deadline" binding:"required"`
	TeacherIDs  []uuid.UUID `json:"teacherIds" binding:"required"`
	FileIDs     []uuid.UUID `json:"fileIds" binding:"required"`
}

type PatchAssignmentRequest struct {
	Title       *string      `json:"title"`
	Description *string      `json:"description"`
	Deadline    *time.Time   `json:"deadline"`
	TeacherIDs  *[]uuid.UUID `json:"teacherIds"`
	FileIDs     *[]uuid.UUID `json:"fileIds"`
}
