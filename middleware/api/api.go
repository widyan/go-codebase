package api

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"codebase/go-codebase/helper"
	"codebase/go-codebase/middleware/interfaces"
	"codebase/go-codebase/middleware/model"
	gmodel "codebase/go-codebase/model"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type ApiImpl struct {
	Rds    *redis.Client
	Logger *logrus.Logger
}

func CreateApi(rds *redis.Client, logger *logrus.Logger) interfaces.ApisMiddleware {
	return &ApiImpl{rds, logger}
}

func (a *ApiImpl) VerifikasiToken(ctx context.Context, token string) (codes int, vrf model.VerifikasiToken, err error) {
	body, err := helper.CallAPI(ctx, a.Logger, os.Getenv("URL_USEETV")+os.Getenv("URL_VERIFIKASI_TOKEN_API_GATEWAY"), "POST", nil, []gmodel.Header{
		{Key: "Authorization", Value: "Bearer " + token},
	})
	if err != nil {
		a.Logger.Error(err.Error())
		return
	}

	// Assign to interface model
	var result map[string]interface{}
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		a.Logger.Error(err.Error())
		return
	}

	if code, ok := result["code"].(float64); ok {
		if code > 399 {
			if errorCode, ok := result["error_code"].(float64); ok {
				codes = int(errorCode)
			}
			err = fmt.Errorf(result["message"].(string))
			a.Logger.Error(err.Error())
			return
		}
	}

	err = json.Unmarshal([]byte(body), &vrf)
	if err != nil {
		a.Logger.Error(err.Error())
		return
	}

	return
}
