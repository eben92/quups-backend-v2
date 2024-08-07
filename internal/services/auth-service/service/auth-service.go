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

	result = mapToUserDTO(u)

	tstring, err := generateAccessToken(local_jwt.AuthContext{
		Sub:  result.ID,
		Name: result.Name,
	})

	if err != nil {
		return result, fmt.Errorf("error signing in. Please try again")
	}

	result.AccessToken = tstring

	return result, nil
}

// AccountSignin handles the user account sign-in process and returns the response user DTO and an error, if any.
func (s *service) AccountSignin(body authdto.AccountSigninDTO) (userdto.TeamMemberDTO, error) {

	authuser, err := local_jwt.GetAuthContext(s.ctx, local_jwt.AUTH_CTX_KEY)

	if err != nil {

		return userdto.TeamMemberDTO{}, fmt.Errorf("error getting account. Please try again")
	}

	uservice := userservice.NewUserService(s.ctx, s.db)
	result, err := uservice.GetUserTeam(body.ID)

	if err != nil {

		return result, fmt.Errorf("error getting account. Please try again")
	}

	tstring, err := generateAccessToken(local_jwt.AuthContext{
		Sub:       authuser.Sub,
		Name:      authuser.Name,
		CompanyID: result.CompanyID,
	})

	if err != nil {
		return result, fmt.Errorf("error signing in. Please try again")
	}

	result.AccessToken = tstring

	return result, nil
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

	result = mapToUserDTO(u)

	tstring, err := generateAccessToken(local_jwt.AuthContext{
		Sub:  result.ID,
		Name: result.Name,
	})

	if err != nil {
		slog.Error("error generating jwt", "Error", err)

		return result, err
	}

	result.AccessToken = tstring

	return result, nil
}

// SoftSignout handles the user sign-out process and returns a string and an error, if any.
// It generates a new jwt token for the user.
func (s *service) SoftSignout() (string, error) {
	user, err := local_jwt.GetAuthContext(s.ctx, local_jwt.AUTH_CTX_KEY)

	if err != nil {

		slog.Error("SoftSignout - error getting user from context", "Error", err)

		return "", fmt.Errorf("error signing out. Please try again")
	}

	tstring, err := generateAccessToken(local_jwt.AuthContext{
		Sub:  user.Sub,
		Name: user.Name,
	})

	if err != nil {
		return "", fmt.Errorf("error signing out. Please try again")
	}

	return tstring, nil
}

func mapToUserDTO(user userdto.UserInternalDTO) authdto.ResponseUserDTO {

	dto := authdto.ResponseUserDTO{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Msisdn:   user.Msisdn,
		ImageUrl: user.ImageUrl,
		Gender:   user.Gender,
	}

	return dto
}

func generateAccessToken(data local_jwt.AuthContext) (string, error) {

	t, err := local_jwt.GenereteJWT(local_jwt.AuthContext{
		Sub:       data.Sub,
		Name:      data.Name,
		CompanyID: data.CompanyID,
	})

	if err != nil {
		slog.Error("error generating jwt", "Error", err)

		return "", err
	}

	tstring := string(t)

	return tstring, nil

}

func isPasswordMatch(rawpass, hashpass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(rawpass))

	return err == nil
}
