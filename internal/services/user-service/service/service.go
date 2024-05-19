package userservice

import (
	"context"
	"quups-backend/internal/database/repository"
)

type Service struct {
	repo *repository.Queries
	ctx  context.Context
}

func New(c context.Context, r *repository.Queries) *Service {
	return &Service{
		repo: r,
		ctx:  c,
	}
}
