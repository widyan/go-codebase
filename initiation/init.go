package initiation

import (
	"codebase/go-codebase/auth"
	"codebase/go-codebase/modules/domain"
	"database/sql"
	"codebase/go-codebase/config"
	"codebase/go-codebase/helper"
	route "codebase/go-codebase/service"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func Init(routesGin *gin.Engine, logger *helper.CustomLogger) (*gin.Engine, *sql.DB, *redis.Client) {
	//bisa nambah config db lain
	db := config.Postgresql(logger) // settingan db postgresql

	//dbRead := config.Postgresql(logger) // settingan dbRead postgresql
	//dbWrite := config.Postgresql(logger) // settingan dbWrite postgresql

	redis := config.Redis(logger)
	response := helper.CreateCustomResponses(os.Getenv("DOMAIN_NAME"))

	//init domain
	userHandler, userUsecase := domain.Init(db, db, redis, logger)
	//init JWT
	initJwt := auth.InitJwt(logger, redis, userUsecase, response)

	routes := route.CreateRoutes(routesGin, initJwt, userHandler)

	return routes.Routes(), db, redis
}
