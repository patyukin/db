package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/patyukin/db/internal/model"
	"github.com/rs/zerolog/log"
)

type Registry struct {
	db *sql.DB
}

//go:generate mockgen -source=registry.go -destination=../../mocks/registry_mock.go -package mocks
type RepositoryInterface interface {
	SelectAllUsers(ctx context.Context) ([]model.User, error)
}

func (r *Registry) GetRepo() RepositoryInterface {
	return &Repository{
		db: r.db,
	}
}

type Handler func(ctx context.Context, repo RepositoryInterface) error

func New(db *sql.DB) *Registry {
	return &Registry{db: db}
}

func (r *Registry) ReadCommitted(ctx context.Context, f Handler) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			if errRollback := tx.Rollback(); errRollback != nil {
				log.Error().Msgf("failed to rollback transaction: %v", errRollback)
			}
		}
	}()

	repo := &Repository{db: tx}

	if err = f(ctx, repo); err != nil {
		return fmt.Errorf("failed to execute handler: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *Registry) Close() error {
	return r.db.Close()
}
