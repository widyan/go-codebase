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
	// u.Cron.Stop()
	c := cron.New(cron.WithParser(cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)))

	ctx := context.Background()
	Result, err := u.Redis.Get(ctx, "worker:lists").Result()
	if err != nil {
		if err.Error() != "redis: nil" {
			return u.Cron
		}
	}

	tasks := []libs.Tasks{}
	json.Unmarshal([]byte(Result), &tasks)
	// log.Println(tasks)
	for _, task := range tasks {
		for _, v := range task.Tasks {
			log.Println("Run task: ", v.Name)
			c.AddFunc(v.Cron, func() {
				log.Println("Run task: ", v.Cron)
				// go u.Rabbit.RunJobs(v.Name)
			})
		}
	}

	c.Start()
	return u.Cron
}
