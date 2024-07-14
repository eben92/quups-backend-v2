package authservice

import (
	"context"
	"quups-backend/internal/database"
	authdto "quups-backend/internal/services/auth-service/dto"
	userdto "quups-backend/internal/services/user-service/dto"
)

const (
	incorrectpass = "incorrect phone number or password"
)

// AuthService represents the interface for the authentication service.
type AuthService interface {
	// Signin handles the user sign-in process and returns the response user DTO and an error, if any.
	Signin(body authdto.SignInRequestDTO) (authdto.ResponseUserDTO, error)

	// Signup handles the user sign-up process and returns the response user DTO and an error, if any.
	Signup(body userdto.CreateUserParams) (authdto.ResponseUserDTO, error)

	// SoftSignout removes companyid from the user's token
	SoftSignout() (string, error)
}

type service struct {
	ctx context.Context
	db  database.Service
}

// NewAuthService creates a new instance of the AuthService interface.
// It takes a context.Context and a database.Service as parameters and returns an AuthService.
func NewAuthService(ctx context.Context, db database.Service) AuthService {
	return &service{
		ctx: ctx,
		db:  db,
	}
}
