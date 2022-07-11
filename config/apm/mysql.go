package config

import (
	"database/sql"
	"time"

	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/mysql"
)

func (c *ConfigImpl) Mysql(dsn, namedb string, SetMaxIdleConns, SetMaxOpenConns int) *sql.DB {
	sqlDB, err := apmsql.Open("mysql", dsn)
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
