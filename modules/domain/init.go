package domain

import (
	"database/sql"

	"github.com/widyan/go-codebase/responses"

	"github.com/widyan/go-codebase/middleware/interfaces"
	"github.com/widyan/go-codebase/modules/domain/handler"
	"github.com/widyan/go-codebase/modules/domain/repository"
	"github.com/widyan/go-codebase/modules/domain/usecase"
	validate "github.com/widyan/go-codebase/validator"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Init(routesGin *gin.Engine, logger *logrus.Logger, validator validate.ValidatorInterface, pq *sql.DB, cfgRseponses responses.GinResponses, authUsecase interfaces.UsecaseInterface) {

	repo := repository.CreateRepository(pq, pq, logger) // Create transaction from db
	userUsecase := usecase.CreateUsecase(repo, logger)

	handler.CreateHandler(userUsecase, logger, cfgRseponses, validator) // Assign function repository for using on handler

	hndler := CreateRoutes(routesGin, authUsecase)

	routesGin = hndler.Routes()

}
