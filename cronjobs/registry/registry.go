package registry

import (
	"codebase/go-codebase/cronjobs/libs"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ interface {
	Worker(task string, job func())
	RunJobs(task string)
}

type RabbitMQImpl struct {
	Conn *amqp.Connection
}

func NewRegister(conn *amqp.Connection) RabbitMQ {
	return &RabbitMQImpl{conn}
}

func (r *RabbitMQImpl) Worker(task string, job func()) {
	ch, err := r.Conn.Channel()
	libs.FailOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		task,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	libs.FailOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	libs.FailOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	libs.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a order: %s", d.Body)
			job()
			log.Printf("Run worker %s success", task)
			d.Ack(false)
		}
	}()

	// log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func (r *RabbitMQImpl) RunJobs(task string) {
	log.Println(task)

	ch, err := r.Conn.Channel()
	libs.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		task,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	libs.FailOnError(err, "Failed to open a channel")

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
	log.Printf(" [x] Run Worker %s", task)
}
