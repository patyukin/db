package db

import (
	"context"
	"github.com/patyukin/db/internal/model"
)

type Repository struct {
	db QueryExecutor
}

func (r *Repository) SelectAllUsers(ctx context.Context) ([]model.User, error) {
	//TODO implement me
	panic("implement me")
}
