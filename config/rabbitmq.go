package config

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *ConfigImpl) RabbitMQ(addrs string) *amqp.Connection {
	conn, err := amqp.Dial(addrs)
	if err != nil {
		c.Logger.Panic(err)
	}

	return conn
}
