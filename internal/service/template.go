package service

import (
	"context"
	"github.com/diianpro/template/internal/storage/mongo"
	"github.com/google/uuid"
)

// Template define service type
type Template struct {
	tmpl *mongo.Storage
}

// New initialize service
func New(tmpl *mongo.Storage) *Template {
	return &Template{
		tmpl: tmpl,
	}
}

// CreateTemplate general template entity
func (t *Template) CreateTemplate(ctx context.Context) error {
	return nil
}

func (t *Template) DeleteTemplate(ctx context.Context, ID uuid.UUID) error {
	return t.tmpl.Delete(ctx, ID)
}
