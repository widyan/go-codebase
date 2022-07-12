package config

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/widyan/go-codebase/config"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)

type ConfigImpl struct {
	Logger *logrus.Logger
}

func CreateConfigImpl(logger *logrus.Logger) config.Config {
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
