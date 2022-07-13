package config

import (
	"context"

	"github.com/go-redis/redis/v8"
	_ "go.elastic.co/apm/module/apmsql/pq"

	apmgoredis "go.elastic.co/apm/module/apmgoredisv8"
)

// Redis is
func (c *ConfigImpl) Redis(address, password string) redis.UniversalClient {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	client.AddHook(apmgoredis.NewHook())

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		c.Logger.Panic(err)
	}

	return client
}
