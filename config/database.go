package config

import "database/sql"

var db *sql.DB

func GetSqlDB() *sql.DB {
	return db
}
