package class

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(router *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	usecase := NewUseCase(repo)
	controller := NewController(usecase)

	RegisterRoutes(router, controller)
}
