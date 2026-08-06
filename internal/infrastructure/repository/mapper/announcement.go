package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"
)

func ToDomainAnnouncement(dbAnnouncement *database.Announcement) *domain.Announcement {
	if dbAnnouncement == nil {
		return nil
	}

	return &domain.Announcement{
		ID:        dbAnnouncement.ID,
		Title:     dbAnnouncement.Title,
		Content:   dbAnnouncement.Content,
		CourseID:  dbAnnouncement.CourseID,
		CreatedAt: dbAnnouncement.CreatedAt,
		UpdatedAt: dbAnnouncement.UpdatedAt,
		Teacher:   ToDomainMaskedUser(&dbAnnouncement.Teacher),
	}
}

func ToDatabaseAnnouncement(announcement *domain.Announcement) *database.Announcement {
	if announcement == nil {
		return nil
	}

	return &database.Announcement{
		ID:        announcement.ID,
		Title:     announcement.Title,
		Content:   announcement.Content,
		CourseID:  announcement.CourseID,
		TeacherID: announcement.Teacher.ID,
		CreatedAt: announcement.CreatedAt,
		UpdatedAt: announcement.UpdatedAt,
	}
}
