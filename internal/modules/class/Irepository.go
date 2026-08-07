package class

import (
	"context"
	classdomain "main/internal/modules/class/domain"

	"github.com/google/uuid"
)

type IRepository interface {
	GetClasses(ctx context.Context) (*[]classdomain.Class, error)
	GetClassByID(ctx context.Context, classID uuid.UUID) (*classdomain.Class, error)
	GetClassBySystemID(ctx context.Context, systemID string) (*classdomain.Class, error)
	CreateClass(ctx context.Context, data *classdomain.Class) (*classdomain.Class, error)
	UpdateClass(ctx context.Context, data *classdomain.Class) (*classdomain.Class, error)
	DeleteClass(ctx context.Context, classID uuid.UUID) error
}
