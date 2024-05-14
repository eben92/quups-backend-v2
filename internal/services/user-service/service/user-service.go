package userservice

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"quups-backend/internal/database/repository"
	model "quups-backend/internal/database/repository"
	userdto "quups-backend/internal/services/user-service/dto"
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

func (s *Service) TestCreate(body *userdto.CreateUserParams) (*model.CreateUserParams, error) {

	if body.Name == nil {
		return nil, fmt.Errorf("user name is required")
	}

	u := &model.CreateUserParams{
		Email: *body.Email,
		Name: sql.NullString{
			String: *body.Name,
		},
	}

	if body.Gender != nil {
		u.Gender.String = *body.Gender
	}

	if body.Msisdn != nil {
		u.Msisdn.String = *body.Msisdn
	}

	if body.Password != nil {

		// todo: hash password here
		u.Password.String = *body.Password
	}

	return u, nil
}

func createUserParams(body *userdto.CreateUserParams) (*model.CreateUserParams, error) {

	if body.Name == nil || body.Email == nil {
		return nil, fmt.Errorf("user name and email is required")
	}

	u := &model.CreateUserParams{
		Email: *body.Email,
		Name: sql.NullString{
			String: *body.Name,
		},
	}

	if body.Gender != nil {
		u.Gender.String = *body.Gender
	}

	if body.Msisdn != nil {
		u.Msisdn.String = *body.Msisdn
	}

	if body.Password != nil {

		// todo: hash password here
		u.Password.String = *body.Password
	}

	return u, nil
}

func (s *Service) Create(body *userdto.CreateUserParams) (*userdto.UserInternalDTO, error) {
	var user *userdto.UserInternalDTO
	params, err := createUserParams(body)

	if err != nil {
		log.Printf("failed to create user error: [%s]", err.Error())

		return nil, err
	}

	log.Printf("about to create new user wih email [%s]", *body.Email)

	u, err := s.repo.CreateUser(s.ctx, *params)

	if err != nil {
		log.Printf("error fetching user with email error:[%s]", err.Error())

		return nil, err
	}

	if u.ID == "" {
		log.Printf("failed to save data in db")

		return nil, fmt.Errorf("failed to create user. Please try again later")
	}

	user = mapToUserInternalDTO(u)

	return user, nil
}

/*
this returns full user dto includinng password
NOTE: response of ths should not be sent to the frontend/client
*/
func (s *Service) FindByEmail(e string) (*userdto.UserInternalDTO, error) {
	log.Printf("fetching user with email [%s]", e)
	var user *userdto.UserInternalDTO

	u, err := s.repo.GetUserByEmail(s.ctx, e)

	if err != nil {
		log.Printf("error fetching user with email [%s] error: [%s]", e, err.Error())

		return nil, err
	}

	if u.ID == "" {
		log.Printf("user with email [%s] does not exist", e)

		return nil, nil
	}

	user = mapToUserInternalDTO(u)

	return user, nil

}

func (s *Service) FindByID(id string) (*userdto.UserInternalDTO, error) {
	log.Printf("fetching user with ID [%s] ", id)
	var user *userdto.UserInternalDTO

	u, err := s.repo.GetUserByID(s.ctx, id)

	if err != nil {
		log.Printf("error fetching user with ID [%s] error: [%s]", id, err.Error())

		return nil, err
	}

	if u.ID == "" {
		log.Printf("user with id [%s] does not exist", id)

		return nil, nil
	}

	user = mapToUserInternalDTO(u)

	return user, nil

}

func (s *Service) FindByMsisdn(msisdn string) (*userdto.UserInternalDTO, error) {
	log.Printf("fetching user with msisdn [%s]\n\n", msisdn)
	var user *userdto.UserInternalDTO

	u, err := s.repo.GetUserByMsisdn(s.ctx, sql.NullString{
		String: msisdn,
		Valid:  true,
	})

	if err != nil {
		log.Printf("error fetching user with msisdn [%s] error: [%s]", msisdn, err.Error())

		return nil, err
	}

	if u.ID == "" {
		log.Printf("user with msisdn [%s] does not exist", msisdn)

		return nil, nil
	}

	user = mapToUserInternalDTO(u)

	return user, nil

}

func mapToUserInternalDTO(user model.User) *userdto.UserInternalDTO {

	dto := &userdto.UserInternalDTO{
		ID:    user.ID,
		Email: user.Email,
	}

	if user.Name.Valid {
		dto.Name = &user.Name.String
	}

	if user.Msisdn.Valid {
		dto.Msisdn = &user.Msisdn.String
	}

	if user.ImageUrl.Valid {
		dto.ImageUrl = &user.ImageUrl.String
	}

	if user.Gender.Valid {
		dto.Gender = &user.Gender.String
	}

	if user.Password.Valid {
		dto.Password = &user.Password.String
	}

	return dto

}
