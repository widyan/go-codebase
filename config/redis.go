package config

import (
	"context"
	"codebase/go-codebase/helper"
	"os"

	"github.com/go-redis/redis/v8"
	apmgoredis "go.elastic.co/apm/module/apmgoredisv8"
)

// Redis is
func Redis(logger *helper.CustomLogger) *redis.Client {
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
