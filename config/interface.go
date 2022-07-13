package config

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConfigImpl struct {
	Logger *logrus.Logger
}

type Config interface {
	Postgresql(dsn string, SetMaxIdleConns, SetMaxOpenConns int) *sql.DB
	Redis(address, password string) redis.UniversalClient
	RabbitMQ(addrs string) *amqp.Connection
	MongoDB(uri, database string) (*mongo.Client, *mongo.Database)
}
