package scheduller

import (
	"codebase/go-codebase/cronjobs/registry"
	"codebase/go-codebase/cronjobs/usecase"
	"codebase/go-codebase/session"
	"context"

	"github.com/sirupsen/logrus"
	amqp "github.com/streadway/amqp"
)

type SchedullerImpl struct {
	Ctx        context.Context
	CronWorker *usecase.CronsWorker
}

func CreateScheduller(connMQ *amqp.Connection, logger *logrus.Logger, project string, session session.Session, Registry registry.RabbitMQ) *SchedullerImpl {
	uscaseWorker := usecase.CreateWorkerClient(logger, project, connMQ, session, Registry)
	return &SchedullerImpl{
		CronWorker: uscaseWorker,
		Ctx:        context.Background(),
	}
	// schdule.InitJob()
}

func (s *SchedullerImpl) TestScheduller() {
	s.CronWorker.Logger.Info("Test Scheduller")
}

func (s *SchedullerImpl) TestLagiAh() {
	s.CronWorker.Logger.Info("Test Lagi Ah")
}

func (s *SchedullerImpl) TestTambahWorker() {
	s.CronWorker.Logger.Info("Tambah Worker Lagi")
}

func (s *SchedullerImpl) TestTambahWorkerLagiDuh() {
	s.CronWorker.Logger.Info("Tambah Worker Lagi Duh")
}

func (s *SchedullerImpl) TestDasarKampret() {
	s.CronWorker.Logger.Info("Dasar Kampret")
}

func (s *SchedullerImpl) TambahWorkerLagi() {
	s.CronWorker.Logger.Info("Tambah Worker Lagi")
}

func (s *SchedullerImpl) RetestWorker() {
	s.CronWorker.Logger.Info("Retest Worker")
}
