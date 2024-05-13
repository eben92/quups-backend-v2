package authservice

import (
	model "quups-backend/internal/database/repository"
)

type userDTO struct {
	ID       string  `json:"id"`
	Email    string  `json:"email"`
	Name     *string `json:"name"`
	Msisdn   *string `json:"msisdn"`
	ImageUrl *string `json:"image_url"`
	Gender   *string `json:"gender"`
}

func mapToUserDTO(user model.User) *userDTO {

	dto := &userDTO{
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

	return dto

}
