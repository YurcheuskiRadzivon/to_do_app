package db

import (
	"database/sql"

	"github.com/YurcheuskiRadzivon/to_do_app/pkg/handlers"
)

type Database interface {
	InitDB(dataSourceName string)
	GetDB() *sql.DB
	CreateTables()
	handlers.AccountHandler
}
