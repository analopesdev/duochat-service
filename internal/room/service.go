package room

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: &repo}
}

func (s *Service) Create(ctx context.Context, r *Room) error {
	return s.repo.Create(ctx, r)
}

func (s *Service) FindAll(ctx context.Context) ([]*Room, error) {
	return s.repo.FindAll(ctx)
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*Room, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID, createdBy uuid.UUID) error {
	return s.repo.Delete(ctx, id, createdBy)
}
