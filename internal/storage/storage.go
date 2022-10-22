package storage

import (
	"context"
	"github.com/google/uuid"
)

// Template database interface
type Template interface {
	Create(ctx context.Context) error
	GetByID(ctx context.Context, Id uuid.UUID) error
	GetByList(ctx context.Context) error
	Delete(ctx context.Context, Id uuid.UUID) error
}
