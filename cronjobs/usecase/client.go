package usecase

import (
	"codebase/go-codebase/cronjobs/registry"
	"context"
	"encoding/json"
	"sync"

	"github.com/go-redis/redis/v8"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type Task struct {
	Name string
	Cron string
}

type Tasks struct {
	Project string
	Tasks   []Task
}

type CronsWorker struct {
	Registry registry.RabbitMQ
	Mutex    sync.Mutex
	Logger   *logrus.Logger
	Redis    *redis.Client
	Project  string
	Task     []Task
}

func CreateWorkerClient(logger *logrus.Logger, redis *redis.Client, project string, connMQ *amqp.Connection) *CronsWorker {
	return &CronsWorker{
		Logger:   logger,
		Redis:    redis,
		Project:  project,
		Registry: registry.NewRegister(connMQ),
	}
}

func (c *CronsWorker) AddJob(ctx context.Context, service, cron string, job func()) {
	c.Mutex.Lock()
	c.Task = append(c.Task, Task{Name: service, Cron: cron})
	c.Mutex.Unlock()

	go c.Registry.Worker(service, job)
}

func (c *CronsWorker) SetListWorker(ctx context.Context) {
	tasks := []Tasks{}

	Result, err := c.Redis.Get(ctx, "worker:lists").Result()
	if err != nil {
		if err.Error() != "redis: nil" {
			c.Logger.Error(err.Error())
			return
		}
	}

	rsltByte := []byte(Result)
	json.Unmarshal(rsltByte, &tasks)

	if len(tasks) == 0 {
		tasks = append(tasks, Tasks{Project: c.Project, Tasks: c.Task})
	} else {
		for idx, v := range tasks {
			if v.Project == c.Project {
				tasks[idx].Tasks = c.Task
			}
		}
	}

	data, err := json.Marshal(tasks)
	if err != nil {
		c.Logger.Error(err.Error())
	}

	if c.Redis.Set(ctx, "worker:lists", data, 0).Err() != nil {
		c.Logger.Error(err.Error())
	}

	if Result != string(data) {
		if c.Redis.Set(ctx, "worker:is_change", 1, 0).Err() != nil {
			return
		}
	}
}
