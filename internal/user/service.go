package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: &repo}
}

func (s *Service) Create(ctx context.Context, u *User) (*User, error) {
	existing, err := s.repo.GetByNickname(ctx, u.Nickname)

	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	if existing != nil {
		fmt.Println("nickname already exists")
		return nil, ErrConflict
	}

	newUser := NewUser(u.Nickname)

	if err := s.repo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *Service) FindAll(ctx context.Context) ([]*User, error) {
	return s.repo.FindAll(ctx)
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) GetByNickname(ctx context.Context, nickname string) (*User, error) {
	return s.repo.GetByNickname(ctx, nickname)
}
