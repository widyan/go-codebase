package config

import (
	"database/sql"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	amqp "github.com/streadway/amqp"
	"go.elastic.co/apm/module/apmsql"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConfigImpl struct {
	Logger *logrus.Logger
}

type Config interface {
	Postgresql(dsn string, SetMaxIdleConns, SetMaxOpenConns int) *sql.DB
	Redis(address, password string) *redis.Client
	RabbitMQ(addrs string) *amqp.Connection
	MongoDB(uri, database string) (*mongo.Client, *mongo.Database)
}

func CreateConfigImplAPM(logger *logrus.Logger) Config {
	return &ConfigImpl{
		Logger: logger,
	}
}

func (c *ConfigImpl) Postgresql(dsn string, SetMaxIdleConns, SetMaxOpenConns int) *sql.DB {
	sqlDB, err := apmsql.Open("postgres", dsn)
	if err != nil {
		c.Logger.Panic(err.Error())
	}
	if err := sqlDB.Ping(); err != nil {
		c.Logger.Panic(err.Error())
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxIdleConns(SetMaxIdleConns)
	sqlDB.SetMaxOpenConns(SetMaxOpenConns)

	return sqlDB
}
