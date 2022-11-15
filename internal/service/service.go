package service

import (
	"context"
	"github.com/diianpro/template/internal/domain"
	"github.com/google/uuid"
)

type Storage interface {
	Create(ctx context.Context, template []byte) (string, error)
	GetByID(ctx context.Context, id uuid.UUID) ([]byte, error)
	GetAll(ctx context.Context, limit int64, offset int64) (*domain.Templates, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// Service define service type
type Service struct {
	tmpl Storage
}

// New initialize service
func New(tmpl Storage) *Service {
	return &Service{
		tmpl: tmpl,
	}
}
