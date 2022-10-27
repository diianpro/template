package storage

import (
	"context"
	"github.com/google/uuid"
	"html/template"
)

// Template database interface
type Template interface {
	Create(ctx context.Context, template []byte) (uuid.UUID, error)
	GetByID(ctx context.Context, Id uuid.UUID) (template.Template, error)
	GetByList(ctx context.Context) error
	Delete(ctx context.Context, Id uuid.UUID) error
}
