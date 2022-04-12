package config

import (
	"codebase/go-codebase/cronjobs/libs"
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	amqp "github.com/rabbitmq/amqp091-go"
)

func RabbitMQ() *amqp.Connection {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ"))
	libs.FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

// Redis is
func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Panic(err)
	}

	return client
}
