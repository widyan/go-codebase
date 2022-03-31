package middleware

import (
	"codebase/go-codebase/helper"
	"codebase/go-codebase/middleware/api"
	"codebase/go-codebase/middleware/config"
	"codebase/go-codebase/middleware/interfaces"
	"codebase/go-codebase/middleware/usecase"
	"os"

	"github.com/sirupsen/logrus"
)

func Init(logger *logrus.Logger, res *helper.Responses) interfaces.UsecaseMiddleware {
	cfg := config.CreateConfig(logger)
	redis := cfg.Redis(os.Getenv("REDIS"), "")

	apis := api.CreateApi(redis, logger)

	return usecase.CreateUsecaseMiddleware(logger, redis, apis, res)
}
