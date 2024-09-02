package usecase

import (
	"context"
	"github.com/patyukin/db/internal/model"
)

type Repository interface {
	SelectAllUsers(ctx context.Context) ([]model.User, error)
}

type UseCase struct {
	repo Repository
}

func New(repo Repository) *UseCase {
	return &UseCase{repo: repo}
}
