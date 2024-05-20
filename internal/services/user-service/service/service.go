package userservice

import (
	"context"
	"errors"
	"quups-backend/internal/database/repository"
)

type Service struct {
	repo *repository.Queries
	ctx  context.Context
}

func New(c context.Context, repo *repository.Queries) *Service {
	return &Service{
		repo: repo,
		ctx:  c,
	}
}

const (
	FOOD    string = "FOOD"
	FASHION string = "FASHION"
)

var BRAND_TYPES = []string{FOOD, FASHION}

var (
	invalidEmailErr     = errors.New("invalid email address.")
	invalidMsisdnErr    = errors.New("invalid phone number.")
	invalidNameErr      = errors.New("name must be greater the 3 characters and excluding any special characters.")
	invalidBrandTypeErr = errors.New("invalid brand type. expecting " + FOOD + " or " + FASHION)
)
