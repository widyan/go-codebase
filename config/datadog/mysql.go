package config

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
	_ "go.elastic.co/apm/module/apmsql/mysql"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)

func (c *ConfigImpl) Mysql(dsn, namedb string, SetMaxIdleConns, SetMaxOpenConns int) *sql.DB {
	sqltrace.Register("mysql", &pq.Driver{}, sqltrace.WithServiceName("metanesia-mysql"))
	sqlDB, err := sqltrace.Open("mysql", dsn)
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
