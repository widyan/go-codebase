package config

import (
	amqp "github.com/streadway/amqp"
)

func (c *ConfigImpl) RabbitMQ(addrs string) *amqp.Connection {
	conn, err := amqp.Dial(addrs)
	if err != nil {
		c.Logger.Panic(err)
	}

	return conn
}
