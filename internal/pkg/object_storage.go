package pkg

import (
	"context"
	"main/internal/domain"
)

type ObjectStorage interface {
	UploadFile(ctx context.Context, bucketName, objectName, contentType string, file []byte) (*domain.File, error)
}

const (
	AssignmentBucket   = "assignments"
	AnnouncementBucket = "announcements"
	MaterialBucket     = "materials"
	SubmissionBucket   = "submissions"
)
