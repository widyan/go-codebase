package config

import (
	helper "codebase/go-codebase/helper/config"
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Cfg helper.Config
}

func CreateConfig(logger *logrus.Logger) *Config {
	return &Config{
		Cfg: helper.CreateGlobalConfig(logger),
	}
}

func (c *Config) Postgresql(dsn, namedb string, SetMaxIdleConns, SetMaxOpenConns int) *sql.DB {
	return c.Cfg.Postgresql(dsn, namedb, SetMaxIdleConns, SetMaxOpenConns)
}

func (c *Config) Redis(address, password string) *redis.Client {
	return c.Cfg.Redis(address, password)
}
