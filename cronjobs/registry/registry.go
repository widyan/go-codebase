package registry

import (
	"codebase/go-codebase/cronjobs/libs"
	"log"

	"github.com/sirupsen/logrus"
	amqp "github.com/streadway/amqp"
)

type RabbitMQ interface {
	Worker(project, task string, job func())
	RunJobs(project, task string)
}

type RabbitMQImpl struct {
	Conn   *amqp.Connection
	Logger *logrus.Logger
}

func NewRegister(conn *amqp.Connection, logger *logrus.Logger) RabbitMQ {
	return &RabbitMQImpl{conn, logger}
}

func (r *RabbitMQImpl) Worker(project, task string, job func()) {
	ch, err := r.Conn.Channel()
	// libs.FailOnError(err, "Failed to open a channel")
	if err != nil {
		r.Logger.Error(err)
		return
	}

	q, err := ch.QueueDeclare(
		project+":"+task, // name
		true,             // durable
		true,             // delete when unused
		false,            // exclusive
		false,            // no-wait
		amqp.Table{
			"x-expires": 1000,
		}, // arguments
	)
	// libs.FailOnError(err, "Failed to declare a queue")
	if err != nil {
		r.Logger.Error(err)
		return
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	// libs.FailOnError(err, "Failed to set QoS")
	if err != nil {
		r.Logger.Error(err)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	// libs.FailOnError(err, "Failed to register a consumer")
	if err != nil {
		r.Logger.Error(err)
		return
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a order from %s with task %s", project, task)
			job()
			log.Printf("Run worker %s with task %s success", project, task)
			d.Ack(false)
		}
	}()

	// log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func (r *RabbitMQImpl) RunJobs(project, task string) {
	ch, err := r.Conn.Channel()
	// libs.FailOnError(err, "Failed to open a channel")
	if err != nil {
		r.Logger.Error(err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		project+":"+task, // name
		true,             // durable
		true,             // delete when unused
		false,            // exclusive
		false,            // no-wait
		amqp.Table{
			"x-expires": 1000,
		}, // arguments
	)
	// libs.FailOnError(err, "Failed to open a channel")
	if err != nil {
		r.Logger.Error(err)
		return
	}

	// body := bodyFrom(os.Args)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(task),
		})
	libs.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Run Worker %s with task %s", project, task)
}
