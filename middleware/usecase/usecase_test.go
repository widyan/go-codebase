package usecase

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/widyan/go-codebase/middleware/entity"
	"github.com/widyan/go-codebase/middleware/interfaces"
	"github.com/widyan/go-codebase/middleware/model"
	iface "github.com/widyan/go-codebase/mocks/middleware/interfaces"
	"github.com/widyan/go-codebase/responses"
)

func CreateTest() (repository *iface.RepositoryInterface, tools *iface.ToolsInterface, usecase interfaces.UsecaseInterface) {
	var logger = logrus.New()
	publicBytes, err := ioutil.ReadFile("../../middleware/public.pem")
	if err != nil {
		logger.Panic(err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		logger.Panic(err)
	}

	privateBytes, err := ioutil.ReadFile("../../middleware/private.pem")
	if err != nil {
		logger.Panic(err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		logger.Panic(err)
	}

	repository = &iface.RepositoryInterface{Mock: mock.Mock{}}
	tools = &iface.ToolsInterface{Mock: mock.Mock{}}
	response := responses.CreateCustomResponses("BACKEND-CHALLENGE")
	usecase = CreateUsecase(repository, tools, response, logger, 1, 1, privateKey, publicKey)
	return
}

var ctx = context.Background()

func TestCreateTokenServices(t *testing.T) {
	var email string = "test@gmail.com"
	repository, tools, usecase := CreateTest()
	users := []entity.User{
		{
			ID:       "df30deb9-4676-444c-91e4-df9eb7fc3a46",
			Email:    email,
			Role:     "admin",
			Name:     "testing",
			IsActive: true,
		}, {
			ID:       "df30deb9-4676-444c-91e4-df9eb7fc3a49",
			Email:    email,
			Role:     "cs",
			Name:     "testing",
			IsActive: true,
		},
	}
	repository.Mock.On("GetUserBasedOnEmail", ctx, email).Return(users, nil)
	var exptoken int64 = 1655489386
	var IssuedAt int64 = 1655474986
	tools.Mock.On("GetTimeNowUnix", 1).Return(exptoken)
	tools.Mock.On("GetTimeNowUnixIssued").Return(IssuedAt)
	request := model.RequestToken{
		Email: email,
	}
	_, err := usecase.CreateTokenServices(ctx, request)
	assert.NoError(t, err)
}

func TestCreateTokenServicesUserNotFound(t *testing.T) {
	var email string = "test@gmail.com"
	repository, tools, usecase := CreateTest()
	users := []entity.User{}
	repository.Mock.On("GetUserBasedOnEmail", ctx, email).Return(users, nil)
	var exptoken int64 = 1655489386
	var IssuedAt int64 = 1655474986
	tools.Mock.On("GetTimeNowUnix", 1).Return(exptoken)
	tools.Mock.On("GetTimeNowUnixIssued").Return(IssuedAt)
	request := model.RequestToken{
		Email: email,
	}
	_, err := usecase.CreateTokenServices(ctx, request)
	assert.Error(t, err)
}

func TestAddUser(t *testing.T) {
	var email string = "test@gmail.com"
	repository, tools, usecase := CreateTest()
	id := "df30deb9-4676-444c-91e4-df9eb7fc3a46"
	tools.Mock.On("GetUUID").Return(id)
	var entityUser = entity.User{
		ID:       id,
		Email:    email,
		Name:     "testing",
		Role:     "admin",
		IsActive: true,
	}
	repository.Mock.On("AddUser", ctx, entityUser).Return(nil)

	var requester = model.RequestUser{
		Email:    email,
		Name:     "testing",
		Role:     "admin",
		IsActive: true,
	}
	err := usecase.AddUser(ctx, requester)
	assert.NoError(t, err)
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestVerifyAutorizationToken(t *testing.T) {
	_, _, usecase := CreateTest()
	r := SetUpRouter()
	r.GET("/", usecase.VerifyAutorizationToken())
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZ21haWwuY29tIiwicm9sZSI6WyJjbGllbnQiLCJkaXJla3R1ciJdLCJleHAiOjE2ODc4NzU2NzIsImlhdCI6MTY1NTQ3NTY3MiwiaXNzIjoiRGFuYSJ9.KPeuSh6ClBeNs8RvbGSIqZ1L-IHVqBQn1BT1fgJQ5D2XSRJPAq_JT6CQ8tmBnhgqZibdt99IYQyvEXjP8wCQ0ZfZmV8kr3DwbUT0_VvO7jqsiOFNMFQmQCoV556Ng7nNaoQziaBI2liCxvvGjF3VkO-A6flSqW7uE594yLM9fslIUiu6Ty972jCld_shWajsnuVPBN-nuI8aBZIYII5_nkycgAiSCRSWuUrrv7DT8hL9tzsz1BLg8y6jAmxWZd99noZ3fnfdmMyVchVDA5LMjPk1zmUA9zZr3dw_foiLaANGdwf5kt7VVGk20o8yP1o2L2Ont71j08_imXmaXSOqgS78lV0OM__1w5CBiP1oI6040-QXfdw4yeaU6ll4_Mx_QHlDHCcNZWj1uAklPaSaYQx-R0S4lG_WfnwUD4ERQqXzqDMwkZoYY21KIxUpX7-HwdMixmE5sWFwXtLSLet27EDZhA9Lxji7kj0KA3WC6MGw6y8Lvo3pDIL6CflIZvsPtCavQ81o7CacHfsE3dnH1EXswoQredjLHdUnZHeaycrYtpSEx3Jk7RoaBrm3ZOJ6X6MAlI6uG9WhwQLSbpNtZsTI_3ckhLeo_GjCHNvlKQSd-EZaR0kg_5at5PYkL7qh5CW0eDKFuPN4LNOYUGQOoIgGUjoYxwvmfM4tyIbVygs")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestVerifyAutorizationTokenWithoutAutorization(t *testing.T) {
	_, _, usecase := CreateTest()
	r := SetUpRouter()
	r.GET("/", usecase.VerifyAutorizationToken())
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestVerifyAutorizationTokenExpired(t *testing.T) {
	_, _, usecase := CreateTest()
	r := SetUpRouter()
	r.GET("/", usecase.VerifyAutorizationToken())
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZ21haWwuY29tIiwicm9sZSI6WyJjbGllbnQiLCJkaXJla3R1ciJdLCJleHAiOjE2NTU0Njk0NjYsImlhdCI6MTY1NTQ1NTA2NiwiaXNzIjoiRGFuYSJ9.sW62nkSDWEnZnCrwsDtxtCVq0do6IkxiahT3xA0v0w4Uhyke9MXs8ZDmmM6VdVeJ20uH_TXrb68Euv1Z9idVX7310HNs17ei74oIfWc5QnaJrBPjN8z5UpQaxxUdrlbglq7wp65RnCSyHtTBXrET4sS9MbMcqZ4yEQULSm8wxtAsAW-DIrN8WCDZzHiN8kMd2USroDijFzM9hQ_mkYffOSTKqD9G_6ZXrAZbSM2P9TzUQOcTX_t9P9X8KptBUq2I_VZzV0HIyAVAVwG6zOeaGde-IBCh7MCj_aAABnrUf4AmrcKX8Ko9ILz2iKYcKnS2jHqt-0u-kaUk29VmZcB8iTApLdl-mjLNATUQ4I6QGTw4D6a1XgT7vTItWqK5wAQT5H-xQrP1E6rurwAKgofVcaAfCz_Y7FQ1k-yvC0qO6_gqK0giB7oP-wlTmDRAwmYqIpKV7qV5e5wIIjMjBKvguqIS80F46Tg8ice8U0CX-I0v5wJMzD9YrAs7rWYwzg3KAdKQftRvWgBNPTQhxZ4NjQdSt02aoU1vzjQwg2-0Q5AWYsKmDPn068ehF9u44lxm7cOk-bOZW4N4o5KFspuMFXFmkwy3yBPT6X5nTWifZC_ojzq_rN6-PAMwMQIYzq4_jadBhon1oEuyMviB14hTMEN5hCCzMivlPu2P5TrgzkY")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestVerifyAutorizationTokenWitoutBearer(t *testing.T) {
	_, _, usecase := CreateTest()
	r := SetUpRouter()
	r.GET("/", usecase.VerifyAutorizationToken())
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZ21haWwuY29tIiwicm9sZSI6WyJjbGllbnQiLCJkaXJla3R1ciJdLCJleHAiOjE2NTU0Njk0NjYsImlhdCI6MTY1NTQ1NTA2NiwiaXNzIjoiRGFuYSJ9.sW62nkSDWEnZnCrwsDtxtCVq0do6IkxiahT3xA0v0w4Uhyke9MXs8ZDmmM6VdVeJ20uH_TXrb68Euv1Z9idVX7310HNs17ei74oIfWc5QnaJrBPjN8z5UpQaxxUdrlbglq7wp65RnCSyHtTBXrET4sS9MbMcqZ4yEQULSm8wxtAsAW-DIrN8WCDZzHiN8kMd2USroDijFzM9hQ_mkYffOSTKqD9G_6ZXrAZbSM2P9TzUQOcTX_t9P9X8KptBUq2I_VZzV0HIyAVAVwG6zOeaGde-IBCh7MCj_aAABnrUf4AmrcKX8Ko9ILz2iKYcKnS2jHqt-0u-kaUk29VmZcB8iTApLdl-mjLNATUQ4I6QGTw4D6a1XgT7vTItWqK5wAQT5H-xQrP1E6rurwAKgofVcaAfCz_Y7FQ1k-yvC0qO6_gqK0giB7oP-wlTmDRAwmYqIpKV7qV5e5wIIjMjBKvguqIS80F46Tg8ice8U0CX-I0v5wJMzD9YrAs7rWYwzg3KAdKQftRvWgBNPTQhxZ4NjQdSt02aoU1vzjQwg2-0Q5AWYsKmDPn068ehF9u44lxm7cOk-bOZW4N4o5KFspuMFXFmkwy3yBPT6X5nTWifZC_ojzq_rN6-PAMwMQIYzq4_jadBhon1oEuyMviB14hTMEN5hCCzMivlPu2P5TrgzkY")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusForbidden, w.Code)
}
