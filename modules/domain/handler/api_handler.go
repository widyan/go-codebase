package handler

import (
	"net/http"
	"strconv"

	"github.com/widyan/go-codebase/helper"
	"github.com/widyan/go-codebase/modules/domain/entity"
	"github.com/widyan/go-codebase/modules/domain/interfaces"
	"github.com/widyan/go-codebase/responses"
	"github.com/widyan/go-codebase/validator"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

type APIHandler struct {
	Usecase   interfaces.Usecase_Interface
	Logger    *logrus.Logger
	Res       responses.GinResponses
	Validator validator.ValidatorInterface
}

var usecase interfaces.Usecase_Interface
var customLogger *logrus.Logger
var response responses.GinResponses
var validate validator.ValidatorInterface

func CreateHandler(Usecase interfaces.Usecase_Interface, logger *logrus.Logger, res responses.GinResponses, vldtr validator.ValidatorInterface) {
	usecase = Usecase
	customLogger = logger
	response = res
	validate = vldtr
}

func GetHandler() *APIHandler {
	return &APIHandler{usecase, customLogger, response, validate}
}

func (a *APIHandler) Test(c *gin.Context) {
	/*
		var User middlemodel.VerifikasiToken
		bind, ok := c.MustGet("bind").([]byte)
		if !ok {
			a.Res.JsonWithErrorCode(c, http.StatusBadRequest, responses.ErrorKetikaMendapatkanDataUser)
			return
		}

		json.Unmarshal(bind, &User)
		json.NewEncoder(c.Writer).Encode(User.Data)
	*/
	a.Res.Json(c, http.StatusOK, map[string]interface{}{
		"key": "value",
	}, "testing")
}

type Booking struct {
	CheckIn  string `form:"check_in" validate:"required"`
	CheckOut string `form:"check_out" validate:"required"`
}

func (a *APIHandler) TestingForm(c *gin.Context) {

	var request Booking

	if err := c.ShouldBindWith(&request, binding.Query); err != nil {
		a.Res.Json(c, http.StatusBadRequest, nil, err.Error())
		return
	}

	err := a.Validator.ValidateRequest(request)
	if err != nil {
		a.Res.Json(c, http.StatusBadRequest, nil, err.Error())
		return
	}

	if err := a.Validator.ValidateRequest(request); err != nil {
		a.Res.Json(c, http.StatusBadRequest, nil, err.Error())
		return
	}

	a.Res.Json(c, http.StatusOK, request, "Success")
}

func (a *APIHandler) InsertUser(c *gin.Context) {
	var param entity.Users

	err := a.Validator.ValidateRequestWithGetBody(c, param)
	if err != nil {
		a.Res.Json(c, http.StatusBadRequest, nil, err.Error())
		return
	}
	err = a.Usecase.InsertUser(c.Request.Context(), param)
	if err != nil {
		a.Res.Json(c, http.StatusInternalServerError, nil, err.Error())
		return
	}

	a.Res.Json(c, http.StatusCreated, param, "Success")
}

func (a *APIHandler) GetOneUser(c *gin.Context) {
	usr, err := a.Usecase.GetOneUser(c.Request.Context())
	if err != nil {
		a.Res.Json(c, http.StatusInternalServerError, nil, err.Error())
		return
	}

	timestamp, err := helper.ConvertTzToNormal(usr.CreatedAt)
	if err != nil {
		a.Res.Json(c, http.StatusInternalServerError, nil, err.Error())
		return
	}

	usr.CreatedAt = timestamp.Format("2006-01-02 15:04:05")

	a.Res.Json(c, http.StatusOK, usr, "Success")
}

func (a *APIHandler) GetAllUsers(c *gin.Context) {
	users, err := a.Usecase.GetAllUsers(c.Request.Context())
	if err != nil {
		a.Res.Json(c, http.StatusInternalServerError, nil, err.Error())
		return
	}

	a.Res.Json(c, http.StatusOK, users, "Success")
}

func (a *APIHandler) UpdateFullnameUserByID(c *gin.Context) {
	var param entity.Users
	if err := c.ShouldBindJSON(&param); err != nil {
		a.Logger.Error(err.Error())
		a.Res.JsonWithErrorCode(c, http.StatusBadRequest, responses.ParameterBodyTidakSesuai)
		return
	}

	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		a.Res.Json(c, http.StatusInternalServerError, nil, err.Error())
		return
	}

	if id == 0 {
		a.Res.JsonWithErrorCode(c, http.StatusBadRequest, responses.IdTidakBoleh0)
		return
	}

	err = a.Usecase.UpdateUserByID(c.Request.Context(), id, param.Fullname)
	if err != nil {
		a.Res.Json(c, http.StatusBadGateway, nil, err.Error())
		return
	}

	a.Res.Json(c, http.StatusOK, nil, "Success")
}

func (a *APIHandler) GetOneUserByID(c *gin.Context) {
	// var User middlemodel.VerifikasiToken
	// bind, ok := c.MustGet("bind").([]byte)
	// if !ok {
	// 	a.Res.JsonWithErrorCode(c, http.StatusBadRequest, helper.ErrorKetikaMendapatkanDataUser)
	// 	return
	// }

	// json.Unmarshal(bind, &User)
	// a.Logger.Println(User)

	id := c.Param("id")
	ids, _ := strconv.Atoi(id)
	users, err := a.Usecase.GetOneUserByID(c.Request.Context(), ids)
	if err != nil {
		a.Res.Json(c, http.StatusInternalServerError, nil, err.Error())
		return
	}

	// users.Fullname = User.Data.Fullname

	a.Res.Json(c, http.StatusOK, users, "Success")
}
