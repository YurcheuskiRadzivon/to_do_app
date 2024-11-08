package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/YurcheuskiRadzivon/to_do_app/pkg/user"
	"github.com/golang-jwt/jwt"
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

func (ps *PostgreSQL) CreateUser(w http.ResponseWriter, req *http.Request) {

	var regReq user.RegisterRequest
	//
	if err := json.NewDecoder(req.Body).Decode(&regReq); err != nil {
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
		http.Error(w, fmt.Sprintf(`{"error": "Error creating user %v"}`, err), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "User registered successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
func (ps *PostgreSQL) LoginUser(w http.ResponseWriter, req *http.Request) {
	var (
		logReq user.LoginRequest
	//errBadCredentials = errors.New("email or password is incorrect")
	)

	if err := json.NewDecoder(req.Body).Decode(&logReq); err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}
	query := fmt.Sprintf(`SELECT password FROM "user" WHERE email = '%s';`, logReq.Email)
	var pass []byte
	err := ps.db.QueryRow(query).Scan(&pass)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, `{"error": "User not found"}`, http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err), http.StatusInternalServerError)
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword(pass, []byte(logReq.Password)); err != nil {
		http.Error(w, `{"error": "Invalid password"}`, http.StatusInternalServerError)
		return
	}
	query = fmt.Sprintf(`SELECT id, username FROM "user" WHERE email = '%s';`, logReq.Email)
	var userId int
	var username string
	if err = ps.db.QueryRow(query).Scan(&userId, &username); err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err), http.StatusNotFound)
		return
	}
	payload := jwt.MapClaims{

		"email":  logReq.Email,
		"name":   username,
		"sub_id": userId,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString(user.GetJwtSecretKey())
	if err != nil {
		http.Error(w, `{"error": "JWT token signing"}`, http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(1 * time.Hour),
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	})
	//response := user.LoginResponse{AccessToken: t}

	response := map[string]string{"message": "Login successful"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)

}
