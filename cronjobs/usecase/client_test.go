package usecase

import (
	"codebase/go-codebase/cronjobs/registry/mocks"
	smocks "codebase/go-codebase/session/mocks"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
)

type MyMockedObject struct {
	mock.Mock
}

func TestAddJob(t *testing.T) {
	mockregistry := new(mocks.RabbitMQ)
	mockSession := new(smocks.Session)
	mockregistry.Mock.On("Worker", "Test-Scheduler-Lagi", "MockTestaja", mock.AnythingOfType("func()"))
	init := CreateWorkerClient(logrus.New(), "Test-Scheduler-Lagi", nil, mockSession, mockregistry)
	init.AddJob("*/1 * * * *", MockTestaja)
}

func MockTestaja() {
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

func TestSetListWorkerIfResultmpty(t *testing.T) {
	var result []byte
	var err error
	result, _ = json.Marshal([]Tasks{})

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

func TestSetListWorkerGetListsError(t *testing.T) {
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
	mockSession.On("Get", ctx, "worker:lists").Return(result, fmt.Errorf("error"))
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

func TestSetListWorkerSetListsError(t *testing.T) {
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
	mockSession.On("Set", ctx, "worker:lists", dataBytes).Return(fmt.Errorf("error"))
	mockSession.On("Set", ctx, "worker:is_change", []byte("1")).Return(err)
	init := CreateWorkerClient(logrus.New(), "Test-Scheduler-Lagi", nil, mockSession, mockregistry)
	init.Task = dataChange

	init.SetListWorker(ctx)
}

func TestSetListWorkerSetIsChangeError(t *testing.T) {
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
	mockSession.On("Set", ctx, "worker:is_change", []byte("1")).Return(fmt.Errorf("error"))
	init := CreateWorkerClient(logrus.New(), "Test-Scheduler-Lagi", nil, mockSession, mockregistry)
	init.Task = dataChange

	init.SetListWorker(ctx)
}
