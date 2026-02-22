package service

import (
	"context"

	"githukudenis.com/db-test-pooling/internal/store"
)

type UserService struct {
	repo store.UserRepository
}

func NewUserService(r store.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetAll(ctx context.Context) ([]*store.User, error) {
	return s.repo.GetAll(ctx)
}
func (s *UserService) GetById(ctx context.Context, id int) (*store.User, error) {
	return s.repo.GetById(ctx, id)
}
func (s *UserService) Save(ctx context.Context, u *store.User) (*store.User, error) {
	return s.repo.Save(ctx, u)
}
