package announcement

import "github.com/google/uuid"

type CreateAnnouncementRequest struct {
	Title    string      `json:"title" binding:"required"`
	Content  string      `json:"content" binding:"required"`
	CourseID uuid.UUID   `json:"courseId" binding:"required,uuid"`
	FileIDs  []uuid.UUID `json:"fileIds" binding:"required"`
}
type PatchAnnouncementRequest struct {
	Title    *string      `json:"title"`
	Content  *string      `json:"content"`
	CourseID *uuid.UUID   `json:"courseId"`
	FileIDs  *[]uuid.UUID `json:"fileIds"`
}
