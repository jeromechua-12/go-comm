package auth

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)

var (
	ErrDuplicateEmail = errors.New("user: duplicate email")
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Insert(ctx context.Context, name, email, password string) error {
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
