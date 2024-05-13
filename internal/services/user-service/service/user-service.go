package userservice

import (
	"context"
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

/*
this returns full user dto includinng password
NOTE: response of ths should not be sent to the frontend/client
*/
func (s *Service) GetUserByEmail(e string) (*userdto.UserInternalDTO, error) {
	log.Printf("fetching user with email %s\n\n", e)
	var user *userdto.UserInternalDTO

	u, err := s.repo.GetUserByEmail(s.ctx, e)

	if err != nil {
		log.Printf("error fetching user with email %s\n error:%s", e, err.Error())

		return nil, err
	}

	if u.ID == "" {
		log.Printf("user with email %s does not exist", e)

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
