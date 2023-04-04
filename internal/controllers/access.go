package controllers

import (
	"backend-code-base-template/internal/usecases"
	"backend-code-base-template/version"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccessController struct {
	router   *gin.RouterGroup
	useCases usecases.AccessUseCaseImply
}

// NewUserController
func NewAccessController(router *gin.RouterGroup, accessUseCase usecases.AccessUseCaseImply) *AccessController {
	return &AccessController{
		router:   router,
		useCases: accessUseCase,
	}
}

// InitRoutes
func (access *AccessController) InitRoutes() {
	access.router.GET("/:version/health", func(ctx *gin.Context) {
		version.RenderHandler(ctx, access, "HealthHandler")
	})
	access.router.GET("/:version/permissions/roles", func(ctx *gin.Context) {
		version.RenderHandler(ctx, access, "GetRoles")
	})
	access.router.GET("/:version/permissions/roles/:id", func(ctx *gin.Context) {
		version.RenderHandler(ctx, access, "GetRolesById")
	})
	access.router.POST("/:version/permissions/roles", func(ctx *gin.Context) {
		version.RenderHandler(ctx, access, "InsertRoles")
	})
	access.router.DELETE("/:version/permissions/roles/:id", func(ctx *gin.Context) {
		version.RenderHandler(ctx, access, "DeleteRoles")
	})
	access.router.PATCH("/:version/permissions/roles/:id", func(ctx *gin.Context) {
		version.RenderHandler(ctx, access, "UpdateRoles")
	})

}

// HealthHandler
func (role *AccessController) HealthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "server run with base version",
	})
}
