package user

import (
	"context"
	"errors"

	"github.com/analopesdev/duochat-service/internal/auth"
)

type Service struct {
	repo *Repository
	auth *auth.Service
}

func NewService(repo Repository, auth *auth.Service) *Service {
	return &Service{repo: &repo, auth: auth}
}

func (s *Service) Create(ctx context.Context, u *User) (string, error) {
	user, err := s.repo.GetByNickname(ctx, u.Nickname)

	if err != nil {
		return "", err
	}

	if user != nil {
		return "", errors.New("user already exists")
	}

	newUser := NewUser(u.Nickname)
	s.repo.Create(ctx, newUser)

	token, err := s.auth.GenerateToken(newUser.ID.String(), newUser.Nickname)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) FindAll(ctx context.Context) ([]*User, error) {
	return s.repo.FindAll(ctx)
}

func (s *Service) GetByID(ctx context.Context, id int64) (*User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) GetByNickname(ctx context.Context, nickname string) (*User, error) {
	return s.repo.GetByNickname(ctx, nickname)
}
