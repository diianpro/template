package service

import (
	"context"
	"github.com/diianpro/template/internal/domain"
	"github.com/google/uuid"
)

type Template interface {
	CreateTemplate(ctx context.Context, template []byte) (string, error)
	GetByID(ctx context.Context, id uuid.UUID) ([]byte, error)
	GetAll(ctx context.Context, limit int64, offset int64) (*domain.Templates, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// CreateTemplate general template entity
func (s *Service) CreateTemplate(ctx context.Context, template []byte) (string, error) {
	return s.tmpl.Create(ctx, template)
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) ([]byte, error) {
	return s.tmpl.GetByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context, limit int64, offset int64) (*domain.Templates, error) {
	return s.tmpl.GetAll(ctx, limit, offset)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.tmpl.Delete(ctx, id)
}
