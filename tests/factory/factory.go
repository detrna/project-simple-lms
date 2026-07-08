package factory

import (
	"main/internal/infrastructure"

	"gorm.io/gorm"
)

var Infra *infrastructure.Infrastructure
var DB *gorm.DB
