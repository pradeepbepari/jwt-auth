package database

import (
	"database/sql"
	"log"

	mysqlConfig "github.com/go-sql-driver/mysql"
)

func ConnectionDB(cfg mysqlConfig.Config) (*sql.DB, error) {
	connection, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Panic(err)
	}
	return connection, nil
}
