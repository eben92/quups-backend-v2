package authservice

import (
	"context"
	"fmt"

	"quups-backend/internal/database/repository"
	authdto "quups-backend/internal/services/auth-service/dto"
	userdto "quups-backend/internal/services/user-service/dto"
	userservice "quups-backend/internal/services/user-service/service"
)

const (
	emailErr    = "Email is required"
	nameErr     = "Name is required"
	enErr       = "Email and name is required"
	msisdnTaken = "Phone number already in use."
)

type Service struct {
	repo *repository.Queries
	ctx  context.Context
}

func New(ctx context.Context, r *repository.Queries) *Service {
	return &Service{
		repo: r,
		ctx:  ctx,
	}
}

func (s *Service) SigninHandler(body *authdto.SignInRequestDTO) (*authdto.ResponseUserDTO, error) {
	uService := userservice.New(s.ctx, s.repo)

	u, err := uService.FindByMsisdn(body.Msisdn)

	if err != nil {
		return nil, fmt.Errorf("incorrect phone number or password")
	}

	user := mapToUserDTO(u)

	return user, nil
}

func (s *Service) SignupHandler(body *userdto.CreateUserParams) (*authdto.ResponseUserDTO, error) {
	uService := userservice.New(s.ctx, s.repo)

	//create user and generate jwt signed token
	u, err := uService.Create(body)

	// send the signed token in both the request body and append it to the browser cookie
	if err != nil {
		return nil, fmt.Errorf("incorrect phone number or password")
	}

	user := mapToUserDTO(u)

	return user, nil
}

func mapToUserDTO(user *userdto.UserInternalDTO) *authdto.ResponseUserDTO {

	dto := &authdto.ResponseUserDTO{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Msisdn:   user.Msisdn,
		ImageUrl: user.ImageUrl,
		Gender:   user.Gender,
	}

	return dto

}
