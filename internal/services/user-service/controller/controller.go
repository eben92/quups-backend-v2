package usercontroller

import "quups-backend/internal/database/repository"

type Controller struct {
	repo *repository.Queries
}

func New(r *repository.Queries) *Controller {
	return &Controller{
		repo: r,
	}
}
