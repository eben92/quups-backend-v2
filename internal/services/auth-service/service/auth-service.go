package authservice

import (
	"fmt"
	"log/slog"

	"golang.org/x/crypto/bcrypt"

	authdto "quups-backend/internal/services/auth-service/dto"
	userdto "quups-backend/internal/services/user-service/dto"
	userservice "quups-backend/internal/services/user-service/service"
	"quups-backend/internal/utils"
	local_jwt "quups-backend/internal/utils/jwt"
)

func (s *service) Signin(body authdto.SignInRequestDTO) (authdto.ResponseUserDTO, error) {
	result := authdto.ResponseUserDTO{}
	uservice := userservice.NewUserService(s.ctx, s.db)

	msisdn, _ := utils.ParseMsisdn(body.Msisdn)

	u, err := uservice.FindByMsisdn(msisdn)

	if err != nil {
		return result, fmt.Errorf(incorrectpass)
	}

	if !isPasswordMatch(body.Password, u.Password) {
		slog.Error("incorrect password for ", "Errors", body.Msisdn)
		return result, fmt.Errorf(incorrectpass)
	}

	user, _ := mapToUserDTO(u)

	return user, nil
}

func (s *service) Signup(body userdto.CreateUserParams) (authdto.ResponseUserDTO, error) {
	uservice := userservice.NewUserService(s.ctx, s.db)
	result := authdto.ResponseUserDTO{}

	// create user and generate jwt signed token
	u, err := uservice.Create(body)
	// send the signed token in both the request body and append it to the browser cookie
	if err != nil {
		return result, err
	}

	result, err = mapToUserDTO(u)

	if err != nil {
		return result, err
	}

	return result, nil
}

func mapToUserDTO(user userdto.UserInternalDTO) (authdto.ResponseUserDTO, error) {
	result := authdto.ResponseUserDTO{}

	t, err := local_jwt.GenereteJWT(user.ID, user.Name)

	if err != nil {
		slog.Error("error generating jwt", "Error", err)

		return result, err
	}

	tstring := string(t)

	dto := authdto.ResponseUserDTO{
		ID:          user.ID,
		Email:       user.Email,
		Name:        user.Name,
		Msisdn:      user.Msisdn,
		ImageUrl:    user.ImageUrl,
		Gender:      user.Gender,
		AccessToken: tstring,
	}

	return dto, nil
}

func isPasswordMatch(rawpass, hashpass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(rawpass))

	return err == nil
}
