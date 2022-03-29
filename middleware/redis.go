package middleware

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	apmgoredis "go.elastic.co/apm/module/apmgoredisv8"
)

// Redis is
func Redis(logger *logrus.Logger) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	client.AddHook(apmgoredis.NewHook())

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		logger.Panic(err)
	}

	return client
}
