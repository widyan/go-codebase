package config

import (
	"database/sql"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	amqp "github.com/streadway/amqp"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/pq"
)

type ConfigImpl struct {
	Logger *logrus.Logger
}

type Config interface {
	Postgresql(dsn, namedb string, SetMaxIdleConns, SetMaxOpenConns int) *sql.DB
	Redis(address, password string) *redis.Client
	RabbitMQ(addrs string) *amqp.Connection
}

func CreateGlobalConfig(logger *logrus.Logger) *ConfigImpl {
	return &ConfigImpl{
		Logger: logger,
	}
}

func (c *ConfigImpl) Postgresql(dsn, namedb string, SetMaxIdleConns, SetMaxOpenConns int) *sql.DB {
	sqlDB, err := apmsql.Open(namedb, dsn)
	if err != nil {
		c.Logger.Panic(err)
	}
	if err := sqlDB.Ping(); err != nil {
		c.Logger.Panic(err)
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxIdleConns(SetMaxIdleConns)
	sqlDB.SetMaxOpenConns(SetMaxOpenConns)

	return sqlDB
}
