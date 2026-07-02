package template

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(router *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	usecase := NewUseCase(repo)
	handler := NewHandler(usecase)

	RegisterRoutes(router, handler)
}
