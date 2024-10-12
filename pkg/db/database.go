package db

import "database/sql"

type Database interface {
	InitDB(dataSourceName string)
	GetDB() *sql.DB
	CreateTables()
}
