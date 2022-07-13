package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/widyan/go-codebase/middleware/interfaces"
	"github.com/widyan/go-codebase/middleware/model"
	"github.com/widyan/go-codebase/responses"
	"github.com/widyan/go-codebase/validator"
)

type APIHandler struct {
	Usecase   interfaces.UsecaseInterface
	Logger    *logrus.Logger
	Response  responses.GinResponses
	Validator validator.ValidatorInterface
}

var usecase interfaces.UsecaseInterface
var customLogger *logrus.Logger
var response responses.GinResponses
var valid validator.ValidatorInterface

func CreateHandler(Usecase interfaces.UsecaseInterface, logger *logrus.Logger, res responses.GinResponses, validate validator.ValidatorInterface) {
	usecase = Usecase
	customLogger = logger
	response = res
	valid = validate
}

func GetHandler() *APIHandler {
	return &APIHandler{usecase, customLogger, response, valid}
}

func (a *APIHandler) Login(c *gin.Context) {
	var request model.RequestToken
	if err := a.Validator.ValidateRequestWithGetBody(c, &request); err != nil {
		a.Response.Json(c, http.StatusBadRequest, nil, err.Error())
		return
	}

	responses, err := a.Usecase.CreateTokenServices(c.Request.Context(), request)
	if err != nil {
		a.Response.JsonWithCaptureError(c, err)
		return
	}

	a.Response.Json(c, http.StatusOK, responses, "Success")
}

func (a *APIHandler) AddUser(c *gin.Context) {
	var request model.RequestUser
	if err := a.Validator.ValidateRequestWithGetBody(c, &request); err != nil {
		a.Response.Json(c, http.StatusBadRequest, nil, err.Error())
		return
	}

	err := a.Usecase.AddUser(c.Request.Context(), request)
	if err != nil {
		a.Response.Json(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	a.Response.Json(c, http.StatusCreated, nil, "Success")
}
