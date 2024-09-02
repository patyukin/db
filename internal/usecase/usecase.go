package usecase

import (
	"context"
	"github.com/patyukin/db/internal/db"
)

type RegistryInterface interface {
	GetRepo() db.RepositoryInterface
	ReadCommitted(ctx context.Context, f db.Handler) error
}

type UseCase struct {
	registry RegistryInterface
}

func New(registry RegistryInterface) *UseCase {
	return &UseCase{registry: registry}
}
