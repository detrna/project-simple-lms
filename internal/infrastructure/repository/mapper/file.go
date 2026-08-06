package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"
)

func ToDomainFile(dbFile *database.File) *domain.File {
	if dbFile == nil {
		return nil
	}

	return &domain.File{
		ID:          dbFile.ID,
		Name:        dbFile.Name,
		FileURL:     dbFile.FileURL,
		Size:        dbFile.Size,
		ContentType: dbFile.ContentType,
		ParentType:  dbFile.ParentType,
		ParentID:    dbFile.ParentID,
		CreatedAt:   dbFile.CreatedAt,
		UpdatedAt:   dbFile.UpdatedAt,
	}
}

func ToDatabaseFile(file *domain.File) *database.File {
	if file == nil {
		return nil
	}

	return &database.File{
		ID:          file.ID,
		Name:        file.Name,
		FileURL:     file.FileURL,
		Size:        file.Size,
		ContentType: file.ContentType,
		ParentType:  file.ParentType,
		ParentID:    file.ParentID,
		CreatedAt:   file.CreatedAt,
		UpdatedAt:   file.UpdatedAt,
	}
}
