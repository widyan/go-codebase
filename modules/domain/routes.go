package domain

import (
	"os"

	"github.com/widyan/go-codebase/middleware/interfaces"
	"github.com/widyan/go-codebase/modules/domain/handler"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	RoutesGin     *gin.Engine
	DomainHandler *handler.APIHandler
}

func CreateRoutes(routesGin *gin.Engine, authUsecase interfaces.UsecaseInterface) *Handlers {
	return &Handlers{
		RoutesGin:     routesGin,
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

	/*
		jwtAuth.Use(handlers.Jwt.VerifyAutorizationToken()) // Verify Authorization
		{
			jwtAuth.GET("/test", handlers.DomainHandler.Test)
			jwtAuth.POST("/user", handlers.DomainHandler.InsertUser)
			jwtAuth.GET("/user/one", handlers.DomainHandler.GetOneUser)
			jwtAuth.PUT("/user/:id", handlers.DomainHandler.UpdateFullnameUserByID)
			jwtAuth.PUT("/formdata", handlers.DomainHandler.TestingForm)
		}
	*/

	return handlers.RoutesGin
}
