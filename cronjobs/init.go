package cronjobs

import (
	"codebase/go-codebase/cronjobs/config"
	"codebase/go-codebase/cronjobs/registry"
	"codebase/go-codebase/cronjobs/usecase"

	"github.com/go-redis/redis/v8"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/robfig/cron/v3"
)

func Init() (rabbitmq *amqp.Connection, redis *redis.Client, crn *cron.Cron) {
	rabbitmq = config.RabbitMQ()
	redis = config.Redis()
	register := registry.NewRegister(rabbitmq)
	initcron := cron.New(cron.WithParser(cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)))
	usecase := usecase.CreateUsecase(register, redis, initcron)
	crn = usecase.CreateTask()

	return
}
