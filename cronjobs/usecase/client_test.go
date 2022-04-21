package usecase

import (
	"codebase/go-codebase/cronjobs/registry/mocks"
	smocks "codebase/go-codebase/session/mocks"
	"context"
	"encoding/json"
	"log"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestAddJob(t *testing.T) {
	mockregistry := new(mocks.RabbitMQ)
	mockSession := new(smocks.Session)
	mockregistry.Mock.On("Worker", "Test-Scheduler-Lagi", MockTest)
	init := CreateWorkerClient(logrus.New(), "Test-Scheduler-Lagi", nil, mockSession, mockregistry)
	init.AddJob("*/1 * * * *", MockTest)
}

func MockTest() {
	log.Println("Test")
}

func TestSetListWorker(t *testing.T) {
	var result []byte
	var err error
	result, _ = json.Marshal([]Tasks{
		{
			Project: "Test-Scheduler-Lagi",
			Tasks: []Task{
				{
					Name: "TestScheduller",
					Cron: "*/1 * * * *",
				}, {
					Name: "TestLagiAh",
					Cron: "*/1 * * * *",
				}, {
					Name: "TestTambahWorker",
					Cron: "*/1 * * * *",
				}, {
					Name: "TestDasarKampret",
					Cron: "*/1 * * * *",
				},
			},
		},
	})

	dataChange := []Task{
		{
			Name: "TestLagiAh",
			Cron: "*/1 * * * *",
		}, {
			Name: "TestTambahWorker",
			Cron: "*/1 * * * *",
		},
	}
	ctx := context.Background()
	mockSession := new(smocks.Session)
	mockregistry := new(mocks.RabbitMQ)
	mockSession.On("Get", ctx, "worker:lists").Return(result, err)
	datas := []Tasks{
		{
			Project: "Test-Scheduler-Lagi",
			Tasks:   dataChange,
		},
	}

	dataBytes, err := json.Marshal(datas)
	mockSession.On("Set", ctx, "worker:lists", dataBytes).Return(err)
	mockSession.On("Set", ctx, "worker:is_change", []byte("1")).Return(err)
	init := CreateWorkerClient(logrus.New(), "Test-Scheduler-Lagi", nil, mockSession, mockregistry)
	init.Task = dataChange

	init.SetListWorker(ctx)
}
