package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"codebase/go-codebase/helper"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type UsecaseMiddlewareImpl struct {
	Logger  *logrus.Logger
	Rds     *redis.Client
	Usecase ApisMiddleware
	Res     *helper.Responses
}

func CreateUsecaseMiddleware(logger *logrus.Logger, rds *redis.Client, usecase ApisMiddleware, res *helper.Responses) UsecaseMiddleware {
	return &UsecaseMiddlewareImpl{logger, rds, usecase, res}
}

func (a *UsecaseMiddlewareImpl) VerifyAutorizationToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			a.Res.AbortWithStatusJSONAndErrorCode(c, 400, helper.TokenTidakBolehKosong)
			return
		}

		// get authorization token
		tokenString := c.GetHeader("Authorization")
		strArr := strings.Split(tokenString, " ")
		if len(strArr) == 2 {
			tokenString = strArr[1]
		} else {
			a.Res.AbortWithStatusJSONAndErrorCode(c, 403, helper.InvalidToken)
			return
		}

		var vrf VerifikasiToken
		tkns, _ := a.Rds.Get(c.Request.Context(), "gw:token:"+tokenString).Result()
		if tkns == "" {
			code, vrf, err := a.Usecase.VerifikasiToken(c.Request.Context(), tokenString)
			if err != nil {
				a.Logger.Error(err.Error())
				a.Res.AbortWithStatusJSONAndInherited(c, 401, code, nil, err.Error())
				return
			}

			byt, _ := json.Marshal(vrf)
			c.Set("bind", byt)
		} else {
			var data ResponsesRedisVerfikasiToken
			err := json.Unmarshal([]byte(tkns), &data)
			if err != nil {
				a.Logger.Error(err.Error())
				a.Res.Json(c, http.StatusInternalServerError, nil, err.Error())
				return
			}

			vrf.Data.IsSuperuser = data.Verifytoken.IsSuperuser
			vrf.Data.DeviceID = data.Verifytoken.Deviceid
			vrf.Data.IsUseeTvUser = data.Verifytoken.IsUseetvUser
			vrf.Data.IsIndiboxUser = data.Verifytoken.IsIndiboxUser
			vrf.Data.IsIndihomeUser = data.Verifytoken.IsIndihomeUser
			vrf.Data.Email = data.Verifytoken.Email
			vrf.Data.Exp = data.Verifytoken.Exp
			vrf.Data.Fullname = data.Verifytoken.Fullname
			vrf.Data.Iat = data.Verifytoken.Iat
			vrf.Data.Iss = data.Verifytoken.Iss
			vrf.Data.LoginStatus = data.Verifytoken.Loginstatus
			vrf.Data.SubscriberID = data.Verifytoken.SubscriberID
			vrf.Data.UserActive = data.Verifytoken.UserActive
			vrf.Data.UserID = data.Verifytoken.Userid

			byt, _ := json.Marshal(vrf)
			c.Set("bind", byt)
		}
	}
}
