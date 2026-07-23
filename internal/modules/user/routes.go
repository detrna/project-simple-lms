package user

import (
	"main/internal/middleware"
	"main/internal/pkg"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	controller    IController
	tokenProvider pkg.JWTProvider
	logger        pkg.Logger
}

func NewRoutes(c IController, tokenProvider pkg.JWTProvider, logger pkg.Logger) *Routes {
	return &Routes{controller: c, tokenProvider: tokenProvider, logger: logger}
}

func (routes Routes) RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/users")

	router.GET("/:id", routes.controller.GetUserByID)
	router.GET(
		"/me",
		middleware.Authenticate(routes.tokenProvider, routes.logger),
		routes.controller.GetMyAccount,
	)
	router.POST(
		"",
		middleware.Authenticate(routes.tokenProvider, routes.logger),
		middleware.RequiredRole("admin", routes.logger),
		routes.controller.CreateUser,
	)
	router.PATCH(
		"/:id/admin",
		middleware.Authenticate(routes.tokenProvider, routes.logger),
		middleware.RequiredRole("admin", routes.logger),
		routes.controller.AdminUpdateUser,
	)
	router.PATCH(
		"/me",
		middleware.Authenticate(routes.tokenProvider, routes.logger),
		routes.controller.UpdateUser,
	)
	router.DELETE(
		"/:id",
		middleware.Authenticate(routes.tokenProvider, routes.logger),
		middleware.RequiredRole("admin", routes.logger),
		routes.controller.DeleteUser,
	)
}
