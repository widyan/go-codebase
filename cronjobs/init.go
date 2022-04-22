package cronjobs

import (
	"codebase/go-codebase/cronjobs/config"
	"codebase/go-codebase/cronjobs/registry"
	"codebase/go-codebase/cronjobs/usecase"
	"codebase/go-codebase/session"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	amqp "github.com/streadway/amqp"
)

func Init(logger *logrus.Logger) (rabbitmq *amqp.Connection, redis *redis.Client, crn *cron.Cron) {
	cfg := config.CreateConfig(logger)
	rabbitmq = cfg.RabbitMQ(os.Getenv("RABBITMQ"))
	redis = cfg.Redis(os.Getenv("REDIS"), "")
	register := registry.NewRegister(rabbitmq, logger)
	initcron := cron.New(cron.WithParser(cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)))
	sesi := session.NewRedisSessionStoreAdapter(redis, 0)
	usecase := usecase.CreateUsecase(logger, register, initcron, sesi)
	crn = usecase.CreateTask()

	return
}
