package libs

import (
	"context"
	"encoding/json"
	"log"
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
	Mutex   sync.Mutex
	ConnMQ  *amqp.Connection
	Logger  *logrus.Logger
	Redis   *redis.Client
	Project string
	Task    []Task
}

func (c *CronsWorker) AddJob(ctx context.Context, service, cron string, job func()) {
	c.Mutex.Lock()
	c.Task = append(c.Task, Task{Name: service, Cron: cron})
	c.Mutex.Unlock()

	go c.Worker(service, job)
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

func (c *CronsWorker) Worker(task string, job func()) {
	ch, err := c.ConnMQ.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		task,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	FailOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			// dotCount := bytes.Count(d.Body, []byte("."))
			// t := time.Duration(dotCount)
			// time.Sleep(t * time.Second)
			job()
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	c.Logger.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
