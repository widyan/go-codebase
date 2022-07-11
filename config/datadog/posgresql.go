package config

import (
	"database/sql"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	amqp "github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
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

func CreateConfigImplDatadog(logger *logrus.Logger) Config {
	return &ConfigImpl{
		Logger: logger,
	}
}

func (c *ConfigImpl) Postgresql(dsn string, SetMaxIdleConns, SetMaxOpenConns int) *sql.DB {
	sqltrace.Register("postgres", &pq.Driver{}, sqltrace.WithServiceName("metanesia-postgres"))
	sqlDB, err := sqltrace.Open("postgres", dsn)
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
