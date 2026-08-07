package class

import (
	"main/internal/middleware"
	"main/internal/pkg"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	controller   IController
	logger       pkg.Logger
	tokenService pkg.TokenService
}

func NewRoutes(c IController, logger pkg.Logger) *Routes {
	return &Routes{controller: c, logger: logger}
}

func (r Routes) RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/classes")
	router.Use(middleware.Authenticate(r.tokenService, r.logger))

	router.GET("/", r.controller.GetClasses)
	router.GET("/:classId", r.controller.GetClassByID)
	router.GET("/system/:systemId", r.controller.GetClassBySystemID)
	router.POST(
		"",
		middleware.RequiredRole("admin", r.logger),
		r.controller.CreateClass,
	)
	router.PATCH(
		"/:classId",
		middleware.RequiredRole("admin", r.logger),
		r.controller.UpdateClass,
	)
	router.DELETE(
		"/:classId",
		middleware.RequiredRole("admin", r.logger),
		r.controller.DeleteClass,
	)
}
