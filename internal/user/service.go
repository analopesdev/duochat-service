package user

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: &repo}
}

func (s *Service) Create(ctx context.Context, u *User) error {
	return s.repo.Create(ctx, u)
}

func (s *Service) GetByID(ctx context.Context, id int64) (*User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) GetByNickname(ctx context.Context, nickname string) (*User, error) {
	return s.repo.GetByNickname(ctx, nickname)
}
