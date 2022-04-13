package middleware

import (
	"codebase/go-codebase/middleware/api"
	"codebase/go-codebase/middleware/config"
	"codebase/go-codebase/middleware/interfaces"
	"codebase/go-codebase/middleware/usecase"
	"codebase/go-codebase/responses"
	"os"

	"github.com/sirupsen/logrus"
)

func Init(logger *logrus.Logger, res responses.GinResponses) interfaces.UsecaseMiddleware {
	cfg := config.CreateConfig(logger)
	redis := cfg.Redis(os.Getenv("REDIS"), "")

	apis := api.CreateApi(redis, logger)

	return usecase.CreateUsecaseMiddleware(logger, redis, apis, res)
}
