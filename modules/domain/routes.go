package domain

import (
	"codebase/go-codebase/middleware"
	"codebase/go-codebase/modules/domain/handler"
	"os"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	RoutesGin     *gin.Engine
	Jwt           middleware.UsecaseMiddleware
	DomainHandler *handler.APIHandler
}

func CreateRoutes(routesGin *gin.Engine, jwt middleware.UsecaseMiddleware) *Handlers {
	return &Handlers{
		RoutesGin:     routesGin,
		Jwt:           jwt,
		DomainHandler: handler.GetHandler(),
	}
}

// Routes is
func (handlers Handlers) Routes() *gin.Engine {
	version := os.Getenv("VERSION_API")

	// Grouping path api
	jwtAuth := handlers.RoutesGin.Group("/" + version + "/api")

	// api with verifikasi jwt token
	jwtAuth.GET("/userall", handlers.DomainHandler.GetAllUsers)

	jwtAuth.Use(handlers.Jwt.VerifyAutorizationToken()) // Verify Authorization
	{
		jwtAuth.GET("/test", handlers.DomainHandler.Test)
		jwtAuth.POST("/user", handlers.DomainHandler.InsertUser)
		jwtAuth.GET("/user/one", handlers.DomainHandler.GetOneUser)
		jwtAuth.PUT("/user/:id", handlers.DomainHandler.UpdateFullnameUserByID)
	}

	return handlers.RoutesGin
}
