package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"
)

func ToDomainComment(dbComment *database.Comment) *domain.Comment {
	if dbComment == nil {
		return nil
	}

	return &domain.Comment{
		ID:         dbComment.ID,
		UserID:     dbComment.UserID,
		ParentType: dbComment.ParentType,
		ParentID:   dbComment.ParentID,
		Content:    dbComment.Content,
		CreatedAt:  dbComment.CreatedAt,
	}
}

func ToDatabaseComment(comment *domain.Comment) *database.Comment {
	if comment == nil {
		return nil
	}

	return &database.Comment{
		ID:         comment.ID,
		Content:    comment.Content,
		UserID:     comment.UserID,
		ParentType: comment.ParentType,
		ParentID:   comment.ParentID,
		CreatedAt:  comment.CreatedAt,
	}
}
