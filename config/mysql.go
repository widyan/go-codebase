package config

import (
	"codebase/go-codebase/helper"
	"database/sql"
	"go.elastic.co/apm/module/apmsql"
	"os"
	"time"
)

func Mysql(logger *helper.CustomLogger) *sql.DB {
	dsn := os.Getenv("GORM_CONNECTION")
	sqlDB, err := apmsql.Open("mysql", dsn)
	if err != nil {
		logger.Panic(err)
	}
	if err := sqlDB.Ping(); err != nil {
		logger.Panic(err)
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(25)

	return sqlDB
}
