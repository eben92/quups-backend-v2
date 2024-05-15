package authservice

import (
	"context"
	"fmt"

	"quups-backend/internal/database/repository"
	authdto "quups-backend/internal/services/auth-service/dto"
	userdto "quups-backend/internal/services/user-service/dto"
	userservice "quups-backend/internal/services/user-service/service"

	"golang.org/x/crypto/bcrypt"
)

const (
	incorrectpass = "incorrect phone number or password"
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
		return nil, fmt.Errorf(incorrectpass)
	}

	if !isPasswordMatch(body.Password, *u.Password) {
		return nil, fmt.Errorf(incorrectpass)
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
		return nil, err
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

func isPasswordMatch(rawpass, hashpass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(rawpass))

	return err == nil

}

func genereteJWT() ([]byte, error) {

	return nil, nil
}
