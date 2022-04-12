package scheduller

import (
	"codebase/go-codebase/cronjobs/libs"
	"context"

	"github.com/go-redis/redis/v8"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type SchedullerImpl struct {
	Ctx        context.Context
	CronWorker *libs.CronsWorker
}

func CreateScheduller(connMQ *amqp.Connection, logger *logrus.Logger, project string, redis *redis.Client) *amqp.Connection {
	schdule := SchedullerImpl{
		CronWorker: &libs.CronsWorker{
			ConnMQ:  connMQ,
			Project: project,
			Redis:   redis,
			Logger:  logger,
		},
		Ctx: context.Background(),
	}
	schdule.InitJob()
	return connMQ
}

func (s *SchedullerImpl) TestScheduller() {
	s.CronWorker.Logger.Info("Test Scheduller")
}

func (s *SchedullerImpl) TestLagiAh() {
	s.CronWorker.Logger.Info("Test Lagi Ah")
}

func (s *SchedullerImpl) TestTambahWorker() {
	s.CronWorker.Logger.Info("Test Lagi Ah")
}
