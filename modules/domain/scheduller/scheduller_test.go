package scheduller

import (
	"codebase/go-codebase/cronjobs/usecase"
	"codebase/go-codebase/session"
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type TestSchedullerImpl struct {
	Ctx     context.Context
	Service SchedullerImpl
}

func CreateSchedullerTest() *TestSchedullerImpl {
	var logger = logrus.New()
	return &TestSchedullerImpl{
		Ctx: context.Background(),
		Service: SchedullerImpl{
			Ctx:        context.Background(),
			CronWorker: usecase.CreateWorkerClient(logger, "", nil, nil),
		},
	}
}

func TestTestScheduller(t *testing.T) {
	var session session.Session
	CreateScheduller(nil, logrus.New(), "", session)
	s := CreateSchedullerTest()
	s.Service.TestScheduller()
	var err error
	assert.NoError(t, err)
}

func TestTestLagiAh(t *testing.T) {
	s := CreateSchedullerTest()
	s.Service.TestLagiAh()
	var err error
	assert.NoError(t, err)
}

func TestTestTambahWorker(t *testing.T) {
	s := CreateSchedullerTest()
	s.Service.TestTambahWorker()
	var err error
	assert.NoError(t, err)
}

func TestTestTambahWorkerLagiDuh(t *testing.T) {
	s := CreateSchedullerTest()
	s.Service.TestTambahWorkerLagiDuh()
	var err error
	assert.NoError(t, err)
}

func TestTestDasarKampret(t *testing.T) {
	s := CreateSchedullerTest()
	s.Service.TestDasarKampret()
	var err error
	assert.NoError(t, err)
}

func TestTambahWorkerLagi(t *testing.T) {
	s := CreateSchedullerTest()
	s.Service.TambahWorkerLagi()
	var err error
	assert.NoError(t, err)
}

func TestRetestWorker(t *testing.T) {
	s := CreateSchedullerTest()
	s.Service.RetestWorker()
	var err error
	assert.NoError(t, err)
}
