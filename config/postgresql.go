package config

import (
	"database/sql"
	"codebase/go-codebase/helper"
	"os"
	"time"

	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/pq"
)

func Postgresql(logger *helper.CustomLogger) *sql.DB {
	dsn := os.Getenv("GORM_CONNECTION")
	sqlDB, err := apmsql.Open("postgres", dsn)
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
