package room_user

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

func (s *Service) Create(ctx context.Context, r *RoomUser) error {
	return s.repo.Create(ctx, r)
}

func (s *Service) FindAll(ctx context.Context) ([]*RoomUser, error) {
	return s.repo.FindAll(ctx)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
