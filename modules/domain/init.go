package domain

import (
	"codebase/go-codebase/helper"
	"codebase/go-codebase/service/api"
	"codebase/go-codebase/service/tools"
	"codebase/go-codebase/service/worker"
	"database/sql"
	"github.com/go-redis/redis/v8"
	"os"
)

func Init(dbRead, dbWrite *sql.DB, redis *redis.Client, logger *helper.CustomLogger) (*APIHandler, *Usecase) {

	repo := CreateRepository(dbWrite, dbRead, logger) // Create transaction from db
	apis := api.CreateApi(redis, logger)
	workers := worker.CreateWorker(logger)
	tools := tools.CreateTools(logger)
	userUsecase := CreateUsecase(repo, apis, workers, tools, logger)

	response := helper.CreateCustomResponses(os.Getenv("DOMAIN_NAME"))

	userHandlers := CreateHandler(userUsecase, redis, logger, response) // Assign function repository for using on handler

	return userHandlers, userUsecase
}

