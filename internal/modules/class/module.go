package class

import (
	"main/internal/infrastructure/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(router *gin.RouterGroup, db *gorm.DB, repo *repository.ClassRepository) {
	usecase := NewUseCase(repo)
	controller := NewController(usecase)

	RegisterRoutes(router, controller)
}
