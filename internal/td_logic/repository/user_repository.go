package repository

import (
	"context"
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/model"
	dberrors "github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/utils/db_errors"
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
	db  *pgxpool.Pool
	lgr *logger.Logger
}

func NewUserRepository(dsnStr string, lgr *logger.Logger) (UserRepository, error) {
	dsn := fmt.Sprintf(dsnStr)
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
		return nil, err
	}
	return &userRepository{
		db:  db,
		lgr: lgr}, nil
}
func (ur *userRepository) GetUser(nickname, email string) (*model.User, error) {
	var User model.User
	query := `SELECT id,nickname,email FROM "User" WHERE nickname=$1 AND email=$2 `
	err := ur.db.QueryRow(context.Background(), query, nickname, email).Scan(&User.ID, &User.Nickname, &User.Email)
	if err != nil {
		return nil, dberrors.UserError(err)
	}
	return &User, nil
}
func (ur *userRepository) InsertUser(User model.UserHash) error {
	query := `INSERT INTO "User"(nickname,email,password) VALUES($1,$2,$3)`
	_, err := ur.db.Exec(context.Background(), query, User.Nickname, User.Email, User.Password)
	if err != nil {
		return dberrors.UserError(err)
	}
	return nil
}
func (ur *userRepository) UpdateUser(id int, User model.User) error {
	query := `UPDATE "User" SET nickname=$1, email=$2 WHERE id=$3`
	_, err := ur.db.Exec(context.Background(), query, User.Nickname, User.Email, id)
	if err != nil {
		return dberrors.UserError(err)
	}
	return nil
}
func (ur *userRepository) DeleteUser(id int) error {
	query := `DELETE FROM "User" WHERE id=$1`
	_, err := ur.db.Exec(context.Background(), query, id)
	if err != nil {
		return dberrors.UserError(err)
	}
	return nil
}
func (ur *userRepository) GetUserPassword(id int) ([]byte, error) {
	var hashedPassword []byte
	query := `SELECT password FROM "User" WHERE id = $1`
	err := ur.db.QueryRow(context.Background(), query, id).Scan(&hashedPassword)
	if err != nil {
		return nil, dberrors.UserError(err)
	}
	return hashedPassword, nil

}
func (ur *userRepository) UpdateUserPassword(id int, pass []byte) error {
	query := `UPDATE "User" SET password=$1 WHERE id=$2`
	_, err := ur.db.Exec(context.Background(), query, pass, id)
	if err != nil {
		return dberrors.UserError(err)
	}
	return nil
}
