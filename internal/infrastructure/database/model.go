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
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP"`

	SubmissionGrades []SubmissionGrades `gorm:"foreignKey:UserID"`
	Enrollment       []Enrollment       `gorm:"foreignKey:StudentID"`
	Teaches          []Teaches          `gorm:"foreignKey:UserID"`
}

type Course struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SystemID     string    `gorm:"not null;uniqueIndex"`
	Name         string    `gorm:"not null"`
	Credits      int       `gorm:"not null"`
	AcademicYear string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP"`

	Teachers      []User         `gorm:"many2many:course_teachers;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Materials     []Material     `gorm:"foreignKey:CourseID"`
	Assignments   []Assignment   `gorm:"foreignKey:CourseID"`
	Announcements []Announcement `gorm:"foreignKey:CourseID"`
	Enrollment    []Enrollment   `gorm:"foreignKey:CourseID"`
}

type Teaches struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"not null"`
	ClassID   uuid.UUID `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP"`

	User  User  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Class Class `gorm:"foreignKey:ClassID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Class struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	SystemID  string    `gorm:"not null;uniqueIndex"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP"`

	Course Course `gorm:"foreignKey:CourseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Material struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	ClassID     uuid.UUID `gorm:"not null"`
	CourseID    uuid.UUID `gorm:"not null;index"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP"`

	Class  Class  `gorm:"foreignKey:ClassID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Course Course `gorm:"foreignKey:CourseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Assignment struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	ClassID     uuid.UUID `gorm:"not null"`
	CourseID    uuid.UUID `gorm:"not null;index"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Deadline    time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP"`

	Class    Class  `gorm:"foreignKey:ClassID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Course   Course `gorm:"foreignKey:CourseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Teachers []User `gorm:"many2many:assignment_teachers;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SubmissionGrades struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	AssignmentID uuid.UUID `gorm:"not null"`
	UserID       uuid.UUID `gorm:"not null"`
	Grade        float64
	CreatedAt    time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP"`

	User       User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Assignment Assignment `gorm:"foreignKey:AssignmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type JWT struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"not null"`
	Token     string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type File struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string    `gorm:"not null"`
	FileURL     string    `gorm:"not null"`
	ContentType string    `gorm:"not null"`
	Size        int64     `gorm:"not null"`
	Bucket      string    `gorm:"not null;index:idx_files_parent"`
	ParentID    uuid.UUID `gorm:"not null;index:idx_files_parent"`
	CreatedAt   time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP"`
}

type Submission struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	AssignmentID uuid.UUID `gorm:"not null;index"`
	StudentID    uuid.UUID `gorm:"not null;index"`
	Score        float64
	TeacherNote  string
	StudentNote  string
	CreatedAt    time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP"`

	Assignment Assignment `gorm:"foreignKey:AssignmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Student    User       `gorm:"foreignKey:StudentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type AcademicTranscript struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	StudentID uuid.UUID `gorm:"not null;uniqueIndex"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP"`

	Student             User                 `gorm:"foreignKey:StudentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SemesterTranscripts []SemesterTranscript `gorm:"foreignKey:AcademicTranscriptID"`
}

type SemesterTranscript struct {
	ID                   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Period               string    `gorm:"not null"`
	AcademicTranscriptID uuid.UUID `gorm:"not null;index"`
	CreatedAt            time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP"`

	AcademicTranscript AcademicTranscript `gorm:"foreignKey:AcademicTranscriptID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Enrollments        []Enrollment       `gorm:"many2many:semester_transcript_enrollments;"`
}

type Announcement struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CourseID  uuid.UUID `gorm:"not null;index"`
	TeacherID uuid.UUID `gorm:"not null;index"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP"`

	Course  Course `gorm:"foreignKey:CourseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Teacher User   `gorm:"foreignKey:TeacherID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Comment struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Content    string    `gorm:"not null"`
	UserID     uuid.UUID `gorm:"not null;index"`
	ParentType string    `gorm:"not null;index"`
	ParentID   uuid.UUID `gorm:"not null;index"`
	CreatedAt  time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Enrollment struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CourseID     uuid.UUID `gorm:"not null;index"`
	StudentID    uuid.UUID `gorm:"not null;index"`
	ClassID      uuid.UUID `gorm:"not null;index"`
	Status       string    `gorm:"not null"`
	AcademicYear string    `gorm:"not null"`
	TeacherNote  *string
	Score        *int
	LeftAt       *time.Time
	CreatedAt    time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP"`

	Course  Course `gorm:"foreignKey:CourseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Student User   `gorm:"foreignKey:StudentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Class   Class  `gorm:"foreignKey:ClassID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
