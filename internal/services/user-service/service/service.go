package userservice

import (
	"context"
	"errors"

	"quups-backend/internal/database"
	"quups-backend/internal/database/repository"
	userdto "quups-backend/internal/services/user-service/dto"
)

type Service interface {
	CompanyService() CompanyService
	UserService() UserService
}

type service struct {
	db  database.Service
	ctx context.Context
}

// New function
func New(c context.Context, db database.Service) Service {
	return &service{
		db:  db,
		ctx: c,
	}
}

func srv(s *service) *service {
	return &service{
		db:  s.db,
		ctx: s.ctx,
	}
}

type UserService interface {
	GetUserTeams(userId string) ([]*userdto.UserTeamDTO, error)
	CreateUserTeam(userId, companyId string) (*repository.Member, error)
	Create(body *userdto.CreateUserParams) (*userdto.UserInternalDTO, error)
	FindByEmail(e string) (*userdto.UserInternalDTO, error)
	FindByID(id string) (*userdto.UserInternalDTO, error)
	FindByMsisdn(msisdn string) (*userdto.UserInternalDTO, error)
}

// UserService method returns User service interface
func (s *service) UserService() UserService {
	return srv(s)
}

type CompanyService interface {
	CreateCompany(
		body *userdto.CreateCompanyParams,
	) (*userdto.CompanyInternalDTO, error)
	GetAllCompanies() ([]*userdto.CompanyInternalDTO, error)
	GetCompanyByName(name string) (*userdto.CompanyInternalDTO, error)
	GetCompanyByID(name string) (*userdto.CompanyInternalDTO, error)
}

// CompanyService method
func (s *service) CompanyService() CompanyService {
	return srv(s)
}

const (
	FOOD    string = "FOOD"
	FASHION string = "FASHION"
)

var BRAND_TYPES = []string{FOOD, FASHION}

var (
	invalidEmailErr  = errors.New("invalid email address.")
	invalidMsisdnErr = errors.New("invalid phone number.")
	invalidNameErr   = errors.New(
		"name must be greater the 3 characters and excluding any special characters.",
	)
	invalidBrandTypeErr = errors.New("invalid brand type. expecting " + FOOD + " or " + FASHION)
)
