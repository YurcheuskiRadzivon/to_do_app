package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	db *sql.DB
}

func (ps *PostgreSQL) InitDB(dataSourceName string) {
	var err error
	ps.db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	if err = ps.db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("PosgreSQL is successfully connected")

}
func (ps *PostgreSQL) GetDB() *sql.DB {
	return ps.db
}
func (ps *PostgreSQL) CreateTables() {

}
