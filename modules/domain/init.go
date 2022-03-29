package domain

import (
	"codebase/go-codebase/helper"
	"database/sql"
	"os"

	"codebase/go-codebase/middleware"
	"codebase/go-codebase/modules/domain/config"
	"codebase/go-codebase/modules/domain/handler"
	"codebase/go-codebase/modules/domain/repository"
	"codebase/go-codebase/modules/domain/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func Init(routesGin *gin.Engine, logger *logrus.Logger) (*gin.Engine, *sql.DB, *redis.Client) {

	cfg := config.CreateConfig(logger)
	redis := cfg.Redis(os.Getenv("REDIS"), "")
	db := cfg.Postgresql(os.Getenv("GORM_CONNECTION"), "postgres", 20, 20)

	//dbRead := config.Postgresql(logger) // settingan dbRead postgresql
	//dbWrite := config.Postgresql(logger) // settingan dbWrite postgresql

	// redis := rds.Redis(logger)
	response := helper.CreateCustomResponses(os.Getenv("DOMAIN_NAME"))

	repo := repository.CreateRepository(db, db, logger) // Create transaction from db
	userUsecase := usecase.CreateUsecase(repo, logger)

	//init JWT
	// initJwt := auth.InitJwt(logger, redis, userUsecase, response)
	middle := middleware.Init(logger, response)

	validator := validator.New()

	response = helper.CreateCustomResponses(os.Getenv("DOMAIN_NAME"))

	handler.CreateHandler(userUsecase, redis, logger, response, validator) // Assign function repository for using on handler

	hndler := CreateRoutes(routesGin, middle)

	routesGin = hndler.Routes()

	return routesGin, db, redis
}
