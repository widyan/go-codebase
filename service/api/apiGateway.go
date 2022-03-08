package api

import (
	"context"
	"encoding/json"
	"fmt"
	"codebase/go-codebase/helper"
	"codebase/go-codebase/model"
	"os"

	"github.com/go-redis/redis/v8"
)

type Api struct {
	Rds    *redis.Client
	Logger *helper.CustomLogger
}

func CreateApi(rds *redis.Client, logger *helper.CustomLogger) *Api {
	return &Api{rds, logger}
}

func (a *Api) VerifikasiToken(ctx context.Context, token string) (codes int, vrf model.VerifikasiToken, err error) {
	body, err := helper.CallAPI(ctx, a.Logger, os.Getenv("BASE_URL")+os.Getenv("URL_VERIFIKASI_TOKEN_API_GATEWAY"), "POST", nil, []model.Header{
		{Key: "Authorization", Value: "Bearer " + token},
	})
	if err != nil {
		a.Logger.Error("(Service) " + err.Error())
		return
	}

	// Assign to interface model
	var result map[string]interface{}
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		a.Logger.Error("(Service) " + err.Error())
		return
	}

	if code, ok := result["code"].(float64); ok {
		if code > 399 {
			if errorCode, ok := result["error_code"].(float64); ok {
				codes = int(errorCode)
			}
			err = fmt.Errorf(result["message"].(string))
			a.Logger.Error("(Service) " + err.Error())
			return
		}
	}

	err = json.Unmarshal([]byte(body), &vrf)
	if err != nil {
		a.Logger.Error("(Service) " + err.Error())
		return
	}

	return
}
