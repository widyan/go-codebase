package middleware

import (
	"os"

	"github.com/widyan/go-codebase/middleware/handler"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	RoutesGin *gin.Engine
	Handler   *handler.APIHandler
}

func CreateRoutes(routesGin *gin.Engine) *Handlers {
	return &Handlers{
		RoutesGin: routesGin,
		Handler:   handler.GetHandler(),
	}
}

// Routes is
func (handlers Handlers) Routes() *gin.Engine {
	version := os.Getenv("VERSION_API")

	// Grouping path api
	jwtAuth := handlers.RoutesGin.Group("/api/auth/" + version)
	jwtAuth.POST("/login", handlers.Handler.Login)
	jwtAuth.POST("/user", handlers.Handler.AddUser)

	return handlers.RoutesGin
}
