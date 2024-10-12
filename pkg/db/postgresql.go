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
	userTable := `
    CREATE TABLE IF NOT EXISTS "user" (
        id SERIAL PRIMARY KEY,
        username VARCHAR(10) NOT NULL,
        email VARCHAR(319) NOT NULL,
        password BYTEA NOT NULL
    );`

	/*taskTable := `
	  CREATE TABLE IF NOT EXISTS tasks (
	      id SERIAL PRIMARY KEY,
	      title VARCHAR(100) NOT NULL,
	      notes TEXT,
	      priority INT NOT NULL,
	      status VARCHAR(20) NOT NULL,
	      user_id INT REFERENCES users(id)
	  );`*/

	_, err := ps.db.Exec(userTable)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	/* _, err = ps.db.Exec(taskTable)
	   if err != nil {
	       log.Fatalf("Failed to create tasks table: %v", err)
	   }*/

	log.Println("Tables created successfully in PostgreSQL!")
}
