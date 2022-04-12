package usecase

import (
	"codebase/go-codebase/cronjobs/libs"
	"codebase/go-codebase/cronjobs/registry"
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
)

type Usecase struct {
	Rabbit registry.RabbitMQ
	Cron   *cron.Cron
	Redis  *redis.Client
}

func CreateUsecase(rabbit registry.RabbitMQ, redis *redis.Client, cron *cron.Cron) Usecase {
	return Usecase{
		Rabbit: rabbit,
		Redis:  redis,
		Cron:   cron,
	}
}

func (u *Usecase) CreateTask() *cron.Cron {
	u.Cron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)))
	ctx := context.Background()
	Result, err := u.Redis.Get(ctx, "worker:lists").Result()
	if err != nil {
		if err.Error() != "redis: nil" {
			return u.Cron
		}
	}

	tasks := []libs.Tasks{}
	json.Unmarshal([]byte(Result), &tasks)
	for _, task := range tasks {
		for _, value := range task.Tasks {
			name := value.Name
			cron := "*/1 * * * * *"
			u.Cron.AddFunc(cron, func() {
				log.Println("Run task: ", name)
				go u.Rabbit.RunJobs(name)
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
	compare, err := u.Redis.Get(ctx, "worker:is_change").Result()
	if err != nil {
		if err.Error() != "redis: nil" {
			return
		}
	}
	if compare == "1" {
		log.Println("MANSSSSSSSSSSSSSSSSSSSSSSSSSSSSS")
		crns := cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)))
		tasks := []libs.Tasks{}
		lists, err := u.Redis.Get(ctx, "worker:lists").Result()
		if err != nil {
			if err.Error() != "redis: nil" {
				return
			}
		}
		json.Unmarshal([]byte(lists), &tasks)
		for _, task := range tasks {
			for _, value := range task.Tasks {
				name := value.Name
				cron := "*/1 * * * * *"
				crns.AddFunc(cron, func() {
					log.Println("Run task: ", name)
					go u.Rabbit.RunJobs(name)
				})
			}
		}
		crns.AddFunc("*/1 * * * * *", func() {
			u.CompareJobs()
		})
		crns.Start()
		u.Cron.Stop()
		u.Cron = crns

		if u.Redis.Set(ctx, "worker:is_change", 0, 0).Err() != nil {
			return
		}
	}
}
