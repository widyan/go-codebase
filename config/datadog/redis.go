package config

import (
	"context"

	"github.com/go-redis/redis/v8"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis.v8"
)

// Redis is
func (c *ConfigImpl) Redis(address, password string) redis.UniversalClient {
	client := redistrace.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       0,        // use default DB
	}, redistrace.WithServiceName("redis-service"))

	// client.AddHook(apmgoredis.NewHook())

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		c.Logger.Panic(err)
	}

	return client
}
