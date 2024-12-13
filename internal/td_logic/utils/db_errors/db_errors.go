package dberrors

import (
	"fmt"
	"github.com/jackc/pgconn"
)

// UserError returns a user-friendly error message based on the database error.
func UserError(err error) error {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch pgErr.SQLState() {
		case "23505": // Unique violation
			if pgErr.ConstraintName == "User_nickname_key" || pgErr.ConstraintName == "User_email_key" {
				return fmt.Errorf("Username or email already exists")
			}
			return fmt.Errorf("Duplicate entry for %s", pgErr.ConstraintName)
		case "22001":
			return fmt.Errorf("Exceeded the length")
		case "23503":
			return fmt.Errorf("Referential integrity constraint violation")
		case "42P01":
			return fmt.Errorf("Table does not exist")
		case "42703":
			return fmt.Errorf("Column does not exist")
		case "22P02":
			return fmt.Errorf("Invalid input format")
		case "40001":
			return fmt.Errorf("Database serialization failure, please try again")
		case "40002":
			return fmt.Errorf("Database deadlock detected, please try again")
		case "P0001":
			return fmt.Errorf("Database function error")
		default:
			return fmt.Errorf("An error occurred with code %s", pgErr.SQLState())
		}
	}
	return fmt.Errorf("Internal server error")
}
