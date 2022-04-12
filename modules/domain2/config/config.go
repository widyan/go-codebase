package config

import (
	"codebase/go-codebase/config"
	"database/sql"

	"github.com/go-redis/redis/v8"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Cfg config.Config
}

func CreateConfig(logger *logrus.Logger) *Config {
	return &Config{
		Cfg: config.CreateGlobalConfig(logger),
	}
}

func (c *Config) Postgresql(dsn, namedb string, SetMaxIdleConns, SetMaxOpenConns int) *sql.DB {
	return c.Cfg.Postgresql(dsn, namedb, SetMaxIdleConns, SetMaxOpenConns)
}

func (c *Config) Redis(address, password string) *redis.Client {
	return c.Cfg.Redis(address, password)
}

func (c *Config) RabbitMQ(address string) *amqp.Connection {
	return c.Cfg.RabbitMQ(address)
}