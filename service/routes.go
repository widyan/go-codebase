package service

import (
	"codebase/go-codebase/auth"
	"codebase/go-codebase/modules/domain"
	"github.com/gin-gonic/gin"
	"os"
)

type Handlers struct {
	RoutesGin *gin.Engine
	jwt *auth.JWT
	DomainHandler *domain.APIHandler
}

func CreateRoutes(routesGin *gin.Engine, jwt *auth.JWT, hand *domain.APIHandler) *Handlers {
	return &Handlers{
		RoutesGin:routesGin,
		jwt:jwt,
		DomainHandler:hand}
}
// Routes is
func (handlers Handlers) Routes() *gin.Engine {
	version := os.Getenv("VERSION_API")

	// Grouping path api
	jwtAuth := handlers.RoutesGin.Group("/" + version + "/api")

	// api with verifikasi jwt token
	jwtAuth.GET("/userall", handlers.DomainHandler.GetAllUsers)

	jwtAuth.Use(handlers.jwt.VerifyAutorizationToken()) // Verify Authorization
	{
		jwtAuth.GET("/test", handlers.DomainHandler.Test)
		jwtAuth.POST("/user", handlers.DomainHandler.InsertUser)
		jwtAuth.GET("/user/one", handlers.DomainHandler.GetOneUser)
		jwtAuth.PUT("/user/:id", handlers.DomainHandler.UpdateFullnameUserByID)
	}

	return handlers.RoutesGin
}
