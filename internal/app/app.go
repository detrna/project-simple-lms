package app

import (
	"main/internal/config"

	"github.com/gin-gonic/gin"
)

func Initialize() error {
	config.Load()

	LoadDatabase()

	return nil
}

func NewApp() *gin.Engine {
	return SetupRouter()
}
