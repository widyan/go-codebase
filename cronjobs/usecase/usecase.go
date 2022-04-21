package usecase

import (
	"codebase/go-codebase/cronjobs/libs"
	"codebase/go-codebase/cronjobs/registry"
	"codebase/go-codebase/session"
	"context"
	"encoding/json"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type Usecase struct {
	Rabbit  registry.RabbitMQ
	Cron    *cron.Cron
	Logger  *logrus.Logger
	Session session.Session
}

func CreateUsecase(logger *logrus.Logger, rabbit registry.RabbitMQ, cron *cron.Cron, session session.Session) Usecase {
	return Usecase{
		Rabbit:  rabbit,
		Cron:    cron,
		Logger:  logger,
		Session: session,
	}
}

func (u *Usecase) CreateTask() *cron.Cron {
	u.Cron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)))
	ctx := context.Background()

	Result, err := u.Session.Get(ctx, "worker:lists")
	if err != nil {
		return u.Cron
	}

	tasks := []libs.Tasks{}
	if string(Result) == "" {
		Result = []byte("[]")
	}

	json.Unmarshal(Result, &tasks)
	for _, task := range tasks {
		for _, value := range task.Tasks {
			project := task.Project
			name := value.Name
			cron := "*/1 * * * * *"
			u.Cron.AddFunc(cron, func() {
				go u.Rabbit.RunJobs(project, name)
			})
		}
	}

	u.Cron.AddFunc("*/1 * * * * *", func() {
		u.CompareJobs()
	})

	u.Cron.Start()
	return u.Cron
}

func (u *Usecase) CompareJobs() {
	ctx := context.Background()
	compare, err := u.Session.Get(ctx, "worker:is_change")
	if err != nil {
		return
	}

	if string(compare) == "1" {
		crns := cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)))
		tasks := []libs.Tasks{}
		lists, err := u.Session.Get(ctx, "worker:lists")
		if err != nil {
			u.Logger.Error(err)
			return
		}
		json.Unmarshal(lists, &tasks)
		for _, task := range tasks {
			for _, value := range task.Tasks {
				project := task.Project
				name := value.Name
				cron := "*/1 * * * * *"
				crns.AddFunc(cron, func() {
					go u.Rabbit.RunJobs(project, name)
				})
			}
		}
		crns.AddFunc("*/1 * * * * *", func() {
			u.CompareJobs()
		})
		u.Cron.Stop()
		crns.Start()
		u.Cron = crns

		if u.Session.Set(ctx, "worker:is_change", []byte("0")) != nil {
			return
		}
	}
}
