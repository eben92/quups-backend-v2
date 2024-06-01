package authservice

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"quups-backend/internal/database"
	authdto "quups-backend/internal/services/auth-service/dto"
	userdto "quups-backend/internal/services/user-service/dto"
	userservice "quups-backend/internal/services/user-service/service"
	local_jwt "quups-backend/internal/utils/jwt"
)

const (
	incorrectpass = "incorrect phone number or password"
)

type service struct {
	ctx context.Context
	db  database.Service
}

func New(ctx context.Context, db database.Service) *service {
	return &service{
		ctx: ctx,
		db:  db,
	}
}

func (s *service) SigninHandler(body *authdto.SignInRequestDTO) (*authdto.ResponseUserDTO, error) {
	uservice := userservice.New(s.ctx, s.db).UserService()

	u, err := uservice.FindByMsisdn(body.Msisdn)
	if err != nil {
		return nil, fmt.Errorf(incorrectpass)
	}

	if !isPasswordMatch(body.Password, *u.Password) {
		log.Printf("incorrect password for [%s]", body.Msisdn)
		return nil, fmt.Errorf(incorrectpass)
	}

	user, _ := mapToUserDTO(u)

	return user, nil
}

func (s *service) SignupHandler(body *userdto.CreateUserParams) (*authdto.ResponseUserDTO, error) {
	uservice := userservice.New(s.ctx, s.db).UserService()

	// create user and generate jwt signed token
	u, err := uservice.Create(body)
	// send the signed token in both the request body and append it to the browser cookie
	if err != nil {
		return nil, err
	}

	user, _ := mapToUserDTO(u)

	return user, nil
}

func mapToUserDTO(user *userdto.UserInternalDTO) (*authdto.ResponseUserDTO, error) {
	t, err := local_jwt.GenereteJWT(user.ID, *user.Name)
	if err != nil {
		return nil, err
	}

	tstring := string(t)

	dto := &authdto.ResponseUserDTO{
		ID:          user.ID,
		Email:       user.Email,
		Name:        user.Name,
		Msisdn:      user.Msisdn,
		ImageUrl:    user.ImageUrl,
		Gender:      user.Gender,
		AccessToken: &tstring,
	}

	return dto, nil
}

func isPasswordMatch(rawpass, hashpass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(rawpass))

	return err == nil
}

/*
Generetes a signed token and return as byte or nil.
Convert to string before sending to client
*/
// func genereteJWT(ID, name string) ([]byte, error) {

// 	// Create a new token object, specifying signing method and the claims
// 	// you would like it to contain.
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"sub":    ID,
// 		"issuer": "WEB",
// 		"name":   name,
// 		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
// 	})

// 	// Sign and get the complete encoded token as a string using the secret
// 	tokenString, err := token.SignedString([]byte(JWT_SECRET))

// 	if err != nil {
// 		log.Printf("Error signing jwt [%s]", err.Error())

// 		return nil, fmt.Errorf("Something went wrong. Please try again. #2")
// 	}

// 	return []byte(tokenString), nil
// }
