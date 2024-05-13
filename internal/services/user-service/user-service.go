package userservice

import "quups-backend/internal/database/repository"

type Service struct {
	repo *repository.Queries
}

func New(r *repository.Queries) *Service {
	return &Service{
		repo: r,
	}
}

func (u *Service) GetUserByEmail() {}
