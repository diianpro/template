package service

import (
	"context"
	"github.com/diianpro/template/internal/storage/mongo"
	"github.com/google/uuid"
	"html/template"
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
func (t *Template) CreateTemplate(ctx context.Context, template []byte) (uuid.UUID, error) {
	return t.tmpl.Create(ctx, template)
}

func (t *Template) GetByID(ctx context.Context, ID uuid.UUID) (template.Template, error) {
	return t.tmpl.GetById(ctx, ID)
}

func (t *Template) Delete(ctx context.Context, ID uuid.UUID) error {
	return t.tmpl.Delete(ctx, ID)
}
