package service

import (
	"context"
	"github.com/diianpro/template/internal/domain"
	"github.com/google/uuid"
)

type Template interface {
	CreateTemplate(ctx context.Context, template *domain.Template) (uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Template, error)
	GetAll(ctx context.Context, limit int64, offset int64) (*domain.Templates, error)
	Delete(ctx context.Context, id string) error
}

// CreateTemplate general template entity
func (s *Service) CreateTemplate(ctx context.Context, template *domain.Template) (uuid.UUID, error) {
	template.ID = uuid.New().String()
	err := s.tmpl.Create(ctx, template)
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.MustParse(template.ID), nil
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*domain.Template, error) {
	return s.tmpl.GetByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context, limit int64, offset int64) (*domain.Templates, error) {
	return s.tmpl.GetAll(ctx, limit, offset)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.tmpl.Delete(ctx, id)
}
