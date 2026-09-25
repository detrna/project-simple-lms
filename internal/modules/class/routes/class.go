package routes

import (
	"main/internal/middleware"
	"main/internal/pkg"

	"github.com/gin-gonic/gin"
)

type IController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	GetBySystemID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type ClassRoutes struct {
	controller   IController
	logger       pkg.Logger
	tokenService pkg.TokenService
}

func NewClassRoutes(c IController, tokenService pkg.TokenService, logger pkg.Logger) *ClassRoutes {
	return &ClassRoutes{
		controller:   c,
		tokenService: tokenService,
		logger:       logger,
	}
}

func (r ClassRoutes) RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/classes")
	router.Use(middleware.Authenticate(r.tokenService, r.logger))

	router.GET("", middleware.HandlePagination(10, 10, r.logger), r.controller.GetAll)
	router.GET("/:id", r.controller.GetByID)
	router.GET("/system/:id", r.controller.GetBySystemID)
	router.POST(
		"",
		middleware.RequiredRole("admin", r.logger),
		r.controller.Create,
	)
	router.PATCH(
		"/:id",
		middleware.RequiredRole("admin", r.logger),
		r.controller.Update,
	)
	router.DELETE(
		"/:id",
		middleware.RequiredRole("admin", r.logger),
		r.controller.Delete,
	)
}
