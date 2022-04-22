package usecase

import (
	smocks "codebase/go-codebase/session/mocks"
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func TestCreateTask(t *testing.T) {
	mockSession := new(smocks.Session)
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
	ctx := context.Background()
	mockSession.On("Get", ctx, "worker:lists").Return(result, err)
	initcron := cron.New(cron.WithParser(cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)))
	init := CreateUsecase(logrus.New(), nil, initcron, mockSession)
	init.CreateTask()
}

func TestCompareJobs(t *testing.T) {
	mockSession := new(smocks.Session)
	var result []byte
	var err error

	ctx := context.Background()
	mockSession.On("Get", ctx, "worker:is_change").Return([]byte("1"), err)

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
	mockSession.On("Get", ctx, "worker:lists").Return(result, err)
	mockSession.On("Set", ctx, "worker:is_change", []byte("0")).Return(err)

	initcron := cron.New(cron.WithParser(cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)))
	init := CreateUsecase(logrus.New(), nil, initcron, mockSession)
	init.CompareJobs()
}

func TestCompareJobsGetListError(t *testing.T) {
	mockSession := new(smocks.Session)
	var result []byte
	var err error

	ctx := context.Background()
	mockSession.On("Get", ctx, "worker:is_change").Return([]byte("1"), err)

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
	mockSession.On("Get", ctx, "worker:lists").Return(result, fmt.Errorf("error"))
	mockSession.On("Set", ctx, "worker:is_change", []byte("0")).Return(err)

	initcron := cron.New(cron.WithParser(cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)))
	init := CreateUsecase(logrus.New(), nil, initcron, mockSession)
	init.CompareJobs()
}

func TestCompareJobsGetIsChangeError(t *testing.T) {
	mockSession := new(smocks.Session)
	var result []byte
	var err error

	ctx := context.Background()
	mockSession.On("Get", ctx, "worker:is_change").Return([]byte("1"), fmt.Errorf("error"))

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
	mockSession.On("Get", ctx, "worker:lists").Return(result, err)
	mockSession.On("Set", ctx, "worker:is_change", []byte("0")).Return(err)

	initcron := cron.New(cron.WithParser(cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)))
	init := CreateUsecase(logrus.New(), nil, initcron, mockSession)
	init.CompareJobs()
}

func TestCompareJobsSetIsChangeError(t *testing.T) {
	mockSession := new(smocks.Session)
	var result []byte
	var err error

	ctx := context.Background()
	mockSession.On("Get", ctx, "worker:is_change").Return([]byte("1"), err)

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
	mockSession.On("Get", ctx, "worker:lists").Return(result, err)
	mockSession.On("Set", ctx, "worker:is_change", []byte("0")).Return(fmt.Errorf("error"))

	initcron := cron.New(cron.WithParser(cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)))
	init := CreateUsecase(logrus.New(), nil, initcron, mockSession)
	init.CompareJobs()
}
