package auth

import (
	"codebase/go-codebase/helper"
	"codebase/go-codebase/model"
	"codebase/go-codebase/modules/domain"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"os"
	"strings"
)

type JWT struct {
	Logger  *helper.CustomLogger
	Rds     *redis.Client
	Usecase domain.Usecase_Interface
	Res     *helper.Responses
}

func InitJwt(logger *helper.CustomLogger, rds *redis.Client, usecase domain.Usecase_Interface, res *helper.Responses) *JWT {
	return &JWT{logger, rds, usecase, res}
}

func (a *JWT) VerifyAutorizationToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		// get authorization token
		const BEARER_SCHEMA = "Bearer "

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			a.Res.AbortWithStatusJSONAndErrorCode(c, 400, helper.TokenTidakBolehKosong)
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		if tokenString == "" || strings.Trim(tokenString, " ") == "Bearer" {
			a.Logger.Error("Invalid Token!")
			a.Res.AbortWithStatusJSONAndErrorCode(c, 403, helper.FormatTokenTidakBenar)
			return
		}
		var vrf model.VerifikasiToken
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
			var data model.ResponsesRedisVerfikasiToken
			err := json.Unmarshal([]byte(tkns), &data)
			if err != nil {
				a.Logger.Error(err.Error())
				a.Res.Json(c, http.StatusInternalServerError, nil, err.Error())
				return
			}

			vrf.Data.DeviceID = data.Verifytoken.Deviceid
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

func (a *JWT) VerifyTokenUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var User model.VerifikasiToken
		bind, ok := c.MustGet("bindDevice").([]byte)
		if !ok {
			// handle error here...
		}
		json.Unmarshal(bind, &User)

		// get authorization token
		// authHeader := c.GetHeader("Authorization")
		authHeader := c.GetHeader("UserAuthorization")
		authHeader = "Bearer " + authHeader
		code, vrf, err := a.Usecase.VerifikasiToken(c.Request.Context(), authHeader)
		if err != nil {
			a.Logger.Error(err.Error())
			a.Res.AbortWithStatusJSONAndInherited(c, 401, code, nil, err.Error())
			return
		}
		if !vrf.Data.LoginStatus {
			a.Logger.Error("Silahkan login terlebih dahulu")
			a.Res.AbortWithStatusJSONAndErrorCode(c, 400, helper.SilahkanLoginTerlebihDahulu)
			return
		}
		if User.Data.DeviceID != vrf.Data.DeviceID {
			a.Logger.Error("DeviceID tidak sesuai")
			a.Res.AbortWithStatusJSONAndErrorCode(c, 400, helper.DeviceIDTidakSesuai)
			return
		}

		byt, _ := json.Marshal(vrf)
		c.Set("bind", byt)
		c.Next()
	}
}

func (a *JWT) VerifyTokenDevice() gin.HandlerFunc {
	return func(c *gin.Context) {

		// get authorization token
		// authHeader := c.GetHeader("Authorization")
		authHeader := c.GetHeader("DeviceAuthorization")
		code, vrf, err := a.Usecase.VerifikasiToken(c.Request.Context(), authHeader)
		if err != nil {
			a.Logger.Error(err.Error())
			a.Res.AbortWithStatusJSONAndInherited(c, 401, code, nil, err.Error())
			return
		}
		if vrf.Data.LoginStatus {
			a.Logger.Error("Header DeviceAuthorization bukan token device autorization")
			a.Res.AbortWithStatusJSONAndErrorCode(c, 403, helper.HeaderDeviceAuthorizationBukanTokenDeviceAutorization)
			return
		}

		byt, _ := json.Marshal(vrf)
		c.Set("bindDevice", byt)
		c.Next()

	}
}

func (a *JWT) VerifyBasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BASIC_SCHEMA = "Basic "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			a.Res.AbortWithStatusJSONAndErrorCode(c, 400, helper.TokenBasicAuthTidakBolehKosong)
			return
		}

		user := os.Getenv("BASIC_AUTH_USERNAME")
		pass := os.Getenv("BASIC_AUTH_PASSWORD")

		basic := base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))
		basic = BASIC_SCHEMA + basic

		if authHeader != basic {
			a.Res.AbortWithStatusJSONAndErrorCode(c, 400, helper.TokenBasicAuthTidakBenar)
			return
		}
	}
}
