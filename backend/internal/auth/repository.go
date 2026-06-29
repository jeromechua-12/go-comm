package auth

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)

type repository struct {
	db *sql.DB
}

func newRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) insert(ctx context.Context, name, email, password string) error {
	query := `INSERT INTO users (name, email, hashed_password, role, created_at)
	VALUES (?, ?, ?, 'customer', UTC_TIMESTAMP)`

	_, err := r.db.ExecContext(ctx, query, name, email, password)
	if err != nil {
		if mySQLError, ok := errors.AsType[*mysql.MySQLError](err); ok {
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "uq_users_email") {
				return ErrDuplicateEmail
			}
		}
		return err
	}
	return nil
}

// Returns the name of a user with email. If bad credential, return error.
func (r* repository) fetch(ctx context.Context, email string) (*User, error) {
	var user User

	query := `SELECT id, name, email, hashed_password, role, created_at FROM users WHERE email = ?` 

	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.Role, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrBadCredentials
		}
		return nil, err
	}

	return &user, nil
}
