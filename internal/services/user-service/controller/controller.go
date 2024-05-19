package usercontroller

import "quups-backend/internal/database/repository"

type UserController struct {
	repo *repository.Queries
}

func New(r *repository.Queries) *UserController {
	return &UserController{
		repo: r,
	}
}
