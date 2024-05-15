package authservice

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"quups-backend/internal/database/repository"
	authdto "quups-backend/internal/services/auth-service/dto"
	userdto "quups-backend/internal/services/user-service/dto"
	userservice "quups-backend/internal/services/user-service/service"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	incorrectpass = "incorrect phone number or password"
)

var (
	JWT_SECRET = os.Getenv("JWT_SECRET")
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

	user, _ := mapToUserDTO(u)

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

	user, _ := mapToUserDTO(u)

	return user, nil
}

func mapToUserDTO(user *userdto.UserInternalDTO) (*authdto.ResponseUserDTO, error) {

	t, err := genereteJWT(user.ID, *user.Name)
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
func genereteJWT(ID, name string) ([]byte, error) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    ID,
		"issuer": "WEB",
		"name":   name,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(JWT_SECRET))

	if err != nil {
		log.Printf("Error signing jwt [%s]", err.Error())

		return nil, fmt.Errorf("Something went wrong. Please try again. #2")
	}

	return []byte(tokenString), nil
}
