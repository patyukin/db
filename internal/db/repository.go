package db

import (
	"context"
	"database/sql"
	"github.com/patyukin/db/internal/model"
)

type QueryExecutor interface {
	ExecContext(ctx context.Context, q string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, q string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, q string, args ...interface{}) *sql.Row
}

type Repository struct {
	db QueryExecutor
}

func (r *Repository) SelectAllUsers(ctx context.Context) ([]model.User, error) {
	return nil, nil
}
