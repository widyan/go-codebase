package domain2

import (
	"codebase/go-codebase/middleware/interfaces"
	"codebase/go-codebase/modules/domain2/handler"
	"os"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	RoutesGin     *gin.Engine
	Jwt           interfaces.UsecaseMiddleware
	DomainHandler *handler.APIHandler
}

func CreateRoutes(routesGin *gin.Engine, jwt interfaces.UsecaseMiddleware) *Handlers {
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
	jwtAuth.GET("/user/:id", handlers.DomainHandler.GetOneUserByID)

	jwtAuth.Use(handlers.Jwt.VerifyAutorizationToken()) // Verify Authorization
	{
		jwtAuth.GET("/test", handlers.DomainHandler.Test)
		jwtAuth.POST("/user", handlers.DomainHandler.InsertUser)
		jwtAuth.GET("/user/one", handlers.DomainHandler.GetOneUser)
		jwtAuth.PUT("/user/:id", handlers.DomainHandler.UpdateFullnameUserByID)
	}

	return handlers.RoutesGin
}
