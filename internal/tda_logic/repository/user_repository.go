package repository

import (
	"context"
	"fmt"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/model"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository interface {
	GetUser(nickname, email string) (*model.User, error)
	InsertUser(User model.UserHash) error
	UpdateUser(id int, User model.User) error
	DeleteUser(id int) error
	GetUserPassword(id int) ([]byte, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(dsnStr string) (UserRepository, error) {
	dsn := fmt.Sprintf(dsnStr)
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
		return nil, err
	}
	return &userRepository{db: db}, nil
}
func (ur *userRepository) GetUser(nickname, email string) (*model.User, error) {
	var User model.User
	query := `SELECT id,name,nickname,email FROM "User" WHERE nickname=$1 AND email=$2 `
	err := ur.db.QueryRow(context.Background(), query, nickname, email).Scan(&User.ID, &User.Name, &User.Nickname, &User.Email)
	if err != nil {
		return nil, err
	}
	return &User, nil
}
func (ur *userRepository) InsertUser(User model.UserHash) error {
	query := `INSERT INTO "User"(name,nickname,email,password) VALUES($1,$2,$3,$4)`
	_, err := ur.db.Exec(context.Background(), query, User.Name, User.Nickname, User.Email, User.Password)
	if err != nil {
		return err
	}
	return nil
}
func (ur *userRepository) UpdateUser(id int, User model.User) error {
	query := `UPDATE "user" SET name=$1, nickname=$2, email=$3 WHERE id=$4`
	_, err := ur.db.Exec(context.Background(), query, User.Name, User.Nickname, User.Email, User.ID)
	if err != nil {
		return err
	}
	return nil
}
func (ur *userRepository) DeleteUser(id int) error {
	query := `DELETE FROM "user" WHERE id=$1`
	_, err := ur.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
func (ur *userRepository) GetUserPassword(id int) ([]byte, error) {
	var hashedPassword []byte
	query := `SELECT password FROM "User" WHERE id = $1`
	err := ur.db.QueryRow(context.Background(), query, id).Scan(&hashedPassword)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil

}
