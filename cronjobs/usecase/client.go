package usecase

import (
	"codebase/go-codebase/cronjobs/registry"
	"context"
	"encoding/json"
	"reflect"
	"runtime"
	"strings"
	"sync"

	"codebase/go-codebase/session"

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
	Project  string
	Task     []Task
	Session  session.Session
}

func CreateWorkerClient(logger *logrus.Logger, project string, connMQ *amqp.Connection, session session.Session) *CronsWorker {
	return &CronsWorker{
		Logger:   logger,
		Project:  project,
		Registry: registry.NewRegister(connMQ),
		Session:  session,
	}
}

func (c *CronsWorker) AddJob(cron string, job func()) {
	service := strings.ReplaceAll(strings.Split(runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name(), ".")[2], "-fm", "")
	c.Mutex.Lock()
	c.Task = append(c.Task, Task{Name: service, Cron: cron})
	c.Mutex.Unlock()

	go c.Registry.Worker(c.Project, service, job)
}

func (c *CronsWorker) SetListWorker(ctx context.Context) {
	tasks := []Tasks{}
	isNewProject := true

	Result, err := c.Session.Get(ctx, "worker:lists")
	if err != nil {
		c.Logger.Error(err.Error())
		return
	}

	json.Unmarshal(Result, &tasks)

	if c.Task == nil {
		c.Task = []Task{}
	}
	if len(tasks) == 0 {
		tasks = append(tasks, Tasks{Project: c.Project, Tasks: c.Task})
	}

	for idx, v := range tasks {
		if v.Project == c.Project {
			isNewProject = false
			tasks[idx].Tasks = c.Task
			break
		}
	}

	if isNewProject {
		tasks = append(tasks, Tasks{Project: c.Project, Tasks: c.Task})
	}

	data, err := json.Marshal(tasks)
	if err != nil {
		c.Logger.Error(err.Error())
	}

	err = c.Session.Set(ctx, "worker:lists", data)
	if err != nil {
		c.Logger.Error(err.Error())
		return
	}

	if string(Result) != string(data) {
		if c.Session.Set(ctx, "worker:is_change", []byte("1")) != nil {
			c.Logger.Error(err.Error())
		}
	}
}
