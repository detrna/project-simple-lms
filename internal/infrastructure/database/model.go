package database

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SystemID  string    `gorm:"not null;uniqueIndex"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"not null;default:user"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`

	SubmissionFile   []SubmissionFile   `gorm:"foreignKey:UserID"`
	SubmissionGrades []SubmissionGrades `gorm:"foreignKey:UserID"`
	Takes            []Takes            `gorm:"foreignKey:UserID"`
	Teaches          []Teaches          `gorm:"foreignKey:UserID"`
}

type Course struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Credits   int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`
}

type Takes struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"not null"`
	ClassID   uuid.UUID `gorm:"not null"`
	Grade     float64
	CreatedAt time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`

	User  User  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Class Class `gorm:"foreignKey:ClassID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Teaches struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"not null"`
	ClassID   uuid.UUID `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`

	User  User  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Class Class `gorm:"foreignKey:ClassID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Class struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	CourseID  uuid.UUID `gorm:"not null"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`

	Course Course `gorm:"foreignKey:CourseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Material struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	ClassID     uuid.UUID `gorm:"not null"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`

	Class        Class          `gorm:"foreignKey:ClassID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MaterialFile []MaterialFile `gorm:"foreignKey:MaterialID"`
}

type MaterialFile struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	MaterialID uuid.UUID `gorm:"not null"`
	URL        string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`

	Material Material `gorm:"foreignKey:MaterialID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Assignment struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	ClassID     uuid.UUID `gorm:"not null"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Deadline    time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`

	Class          Class            `gorm:"foreignKey:ClassID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AssignmentFile []AssignmentFile `gorm:"foreignKey:AssignmentID"`
	SubmissionFile []SubmissionFile `gorm:"foreignKey:AssignmentID"`
}

type AssignmentFile struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	AssignmentID uuid.UUID `gorm:"not null"`
	URL          string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`

	Assignment Assignment `gorm:"foreignKey:AssignmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SubmissionFile struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	AssignmentID uuid.UUID `gorm:"not null"`
	UserID       uuid.UUID `gorm:"not null"`
	URL          string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`

	User       User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Assignment Assignment `gorm:"foreignKey:AssignmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SubmissionGrades struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	AssignmentID uuid.UUID `gorm:"not null"`
	UserID       uuid.UUID `gorm:"not null"`
	Grade        float64
	CreatedAt    time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`

	User       User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Assignment Assignment `gorm:"foreignKey:AssignmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type JWT struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"not null"`
	Token     string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
