package config

import (
	"codebase/go-codebase/cronjobs/libs"
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	amqp "github.com/rabbitmq/amqp091-go"
)

func RabbitMQ() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	libs.FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

// Redis is
func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.60.160.76:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Panic(err)
	}

	return client
}
