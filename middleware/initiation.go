package middleware

import (
	"codebase/go-codebase/helper"
	"os"

	"github.com/sirupsen/logrus"
)

func Init(logger *logrus.Logger, res *helper.Responses) UsecaseMiddleware {
	cfg := CreateConfig(logger)
	redis := cfg.Redis(os.Getenv("REDIS"), "")

	apis := CreateApi(redis, logger)

	return CreateUsecaseMiddleware(logger, redis, apis, res)
}
