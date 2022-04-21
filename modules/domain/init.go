package domain

import (
	"codebase/go-codebase/responses"
	"codebase/go-codebase/session"
	"database/sql"
	"os"

	"codebase/go-codebase/middleware"
	"codebase/go-codebase/modules/domain/config"
	"codebase/go-codebase/modules/domain/handler"
	"codebase/go-codebase/modules/domain/repository"
	"codebase/go-codebase/modules/domain/scheduller"
	"codebase/go-codebase/modules/domain/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/go-redis/redis/v8"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

func Init(routesGin *gin.Engine, logger *logrus.Logger) (*gin.Engine, *sql.DB, *redis.Client, *amqp091.Connection) {

	cfg := config.CreateConfig(logger)
	redis := cfg.Redis(os.Getenv("REDIS"), "")
	db := cfg.Postgresql(os.Getenv("GORM_CONNECTION"), "postgres", 20, 20)
	connMQ := cfg.RabbitMQ(os.Getenv("RABBITMQ"))

	//dbRead := config.Postgresql(logger) // settingan dbRead postgresql
	//dbWrite := config.Postgresql(logger) // settingan dbWrite postgresql

	// redis := rds.Redis(logger)
	response := responses.CreateCustomResponses(os.Getenv("DOMAIN_NAME"))

	repo := repository.CreateRepository(db, db, logger) // Create transaction from db
	userUsecase := usecase.CreateUsecase(repo, logger)

	//init JWT
	middle := middleware.Init(logger, response)

	validator := validator.New()

	response = responses.CreateCustomResponses(os.Getenv("DOMAIN_NAME"))

	handler.CreateHandler(userUsecase, logger, response, validator) // Assign function repository for using on handler

	sesi := session.NewRedisSessionStoreAdapter(redis, 0)
	initCron := scheduller.CreateScheduller(connMQ, logger, os.Getenv("DOMAIN_NAME"), sesi)
	initCron.InitJob()

	hndler := CreateRoutes(routesGin, middle)

	routesGin = hndler.Routes()

	return routesGin, db, redis, connMQ
}
