package config

import (
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/widyan/go-codebase/config"
	"go.elastic.co/apm/module/apmsql"
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
