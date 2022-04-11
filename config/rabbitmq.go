package config

import (
	"codebase/go-codebase/cronjobs/libs"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *ConfigImpl) RabbitMQ(addrs string) *amqp.Connection {
	conn, err := amqp.Dial(addrs)
	libs.FailOnError(err, "Failed to connect to RabbitMQ")

	return conn
}
