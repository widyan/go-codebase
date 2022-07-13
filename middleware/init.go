package middleware

import (
	"database/sql"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/widyan/go-codebase/middleware/handler"
	"github.com/widyan/go-codebase/middleware/interfaces"
	"github.com/widyan/go-codebase/middleware/repository"
	"github.com/widyan/go-codebase/middleware/tools"
	"github.com/widyan/go-codebase/middleware/usecase"
	"github.com/widyan/go-codebase/responses"
	validate "github.com/widyan/go-codebase/validator"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Init(routesGin *gin.Engine, logger *logrus.Logger, pq *sql.DB, validator validate.ValidatorInterface, response responses.GinResponses) interfaces.UsecaseInterface {
	publicBytes, err := ioutil.ReadFile("./middleware/public.pem")
	if err != nil {
		logger.Panic(err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		logger.Panic(err)
	}

	privateBytes, err := ioutil.ReadFile("./middleware/private.pem")
	if err != nil {
		logger.Panic(err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		logger.Panic(err)
	}

	tool := tools.CreateTools()
	repo := repository.CreateRepository(pq, pq, logger)
	authUsecase := usecase.CreateUsecase(repo, tool, response, logger, 4, 8, privateKey, publicKey)
	handler.CreateHandler(authUsecase, logger, response, validator)
	hndler := CreateRoutes(routesGin)
	routesGin = hndler.Routes()
	return authUsecase
}
