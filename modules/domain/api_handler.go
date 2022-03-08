package domain

import (
	"codebase/go-codebase/entity"
	"codebase/go-codebase/helper"
	"codebase/go-codebase/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strconv"
)

type APIHandler struct {
	Usecase Usecase_Interface
	Rds     *redis.Client
	Logger  *helper.CustomLogger
	Res     *helper.Responses
}

var usecase Usecase_Interface
var rdsClient *redis.Client
var customLogger *helper.CustomLogger
var responses *helper.Responses

func CreateHandler(Usecase Usecase_Interface, rds *redis.Client, logger *helper.CustomLogger, res *helper.Responses) {
	usecase = Usecase
	rdsClient = rds
	customLogger = logger
	responses = res
}

func GetHandler()*APIHandler  {
	return&APIHandler{usecase, rdsClient, customLogger, responses}
}

func (a APIHandler) Test(c *gin.Context) {
	var User model.VerifikasiToken
	bind, ok := c.MustGet("bind").([]byte)
	if !ok {
		a.Res.JsonWithErrorCode(c, http.StatusBadRequest, helper.ErrorKetikaMendapatkanDataUser)
		return
	}

	json.Unmarshal(bind, &User)
	a.Res.Json(c, http.StatusOK, User.Data, "testing")
}

func (a APIHandler) InsertUser(c *gin.Context) {
	var param entity.Users
	if err := c.ShouldBindJSON(&param); err != nil {
		a.Logger.Error(err.Error())
		a.Res.JsonWithErrorCode(c, http.StatusBadRequest, helper.ParameterBodyTidakSesuai)
		return
	}

	a.Logger.ErrorWithContext(c, param, "Test out context")

	if param.Fullname == "" {
		a.Res.JsonWithErrorCode(c, http.StatusBadRequest, helper.FullNameTidakBolehKosong)
		return
	}

	err := a.Usecase.InsertUser(c.Request.Context(), param)
	if err != nil {
		a.Res.Json(c, http.StatusInternalServerError, nil, err.Error())
		return
	}

	a.Res.Json(c, http.StatusCreated, param, "Success")
	return
}

func (a APIHandler) GetOneUser(c *gin.Context) {
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
	return
}

func (a APIHandler) GetAllUsers(c *gin.Context) {
	users, err := a.Usecase.GetAllUsers(c.Request.Context())
	if err != nil {
		a.Res.Json(c, http.StatusInternalServerError, nil, err.Error())
	}

	a.Res.Json(c, http.StatusOK, users, "Success")
	return
}

func (a APIHandler) UpdateFullnameUserByID(c *gin.Context) {
	var param entity.Users
	if err := c.ShouldBindJSON(&param); err != nil {
		a.Logger.Error(err.Error())
		a.Res.JsonWithErrorCode(c, http.StatusBadRequest, helper.ParameterBodyTidakSesuai)
		return
	}

	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		a.Res.Json(c, http.StatusInternalServerError, nil, err.Error())
		return
	}

	if id == 0 {
		a.Res.JsonWithErrorCode(c, http.StatusBadRequest, helper.IdTidakBoleh0)
		return
	}

	err = a.Usecase.UpdateUserByID(c.Request.Context(), id, param.Fullname)
	if err != nil {
		a.Res.Json(c, http.StatusBadGateway, nil, err.Error())
	}

	a.Res.Json(c, http.StatusOK, nil, "Success")
	return
}
