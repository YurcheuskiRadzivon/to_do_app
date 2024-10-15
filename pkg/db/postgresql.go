package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/YurcheuskiRadzivon/to_do_app/pkg/user"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
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
        username VARCHAR(10) UNIQUE NOT NULL,
        email VARCHAR(319)  UNIQUE NOT NULL,
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

func (ps *PostgreSQL) CreateAccount(w http.ResponseWriter, req *http.Request) {

	var regReq user.RegUser
	err := json.NewDecoder(req.Body).Decode(&regReq)
	if err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regReq.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, `{"error": "Error setting password"}`, http.StatusInternalServerError)
		return
	}
	query := `INSERT INTO "user" (username, email, password) VALUES ($1, $2, $3)`
	_, err = ps.db.Exec(query, regReq.Username, regReq.Email, hashedPassword)
	if err != nil {
		http.Error(w, `{"error": "Error creating user"}`, http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "User registered successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
