package userservice

import (
	"context"
	"errors"

	"quups-backend/internal/database"
	"quups-backend/internal/database/repository"
	userdto "quups-backend/internal/services/user-service/dto"
)

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

// Service represents the interface for the user service.
type Service interface {
	CompanyService() CompanyService
	UserService() UserService
	NewPaymentService() PaymentService
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

func serviceProvider(s *service) *service {
	return &service{
		db:  s.db,
		ctx: s.ctx,
	}
}

// UserService represents the interface for user-related operations.
type UserService interface {
	// GetUserTeams retrieves the teams that a user belongs to.
	GetUserTeams(userId string) ([]userdto.UserTeamDTO, error)

	// CreateUserTeam creates a new user team for a given company.
	CreateUserTeam(companyId string) (repository.Member, error)

	// Create creates a new user with the provided parameters.
	Create(body userdto.CreateUserParams) (userdto.UserInternalDTO, error)

	// FindByEmail retrieves a user by their email address.
	FindByEmail(e string) (userdto.UserInternalDTO, error)

	// FindByID retrieves a user by their ID.
	FindByID(id string) (userdto.UserInternalDTO, error)

	// FindByMsisdn retrieves a user by their MSISDN (mobile number).
	FindByMsisdn(msisdn string) (userdto.UserInternalDTO, error)
}

// UserService method returns User service interface
func (s *service) UserService() UserService {
	return serviceProvider(s)
}

// CompanyService represents the interface for managing company-related operations.
type CompanyService interface {
	// CreateCompany creates a new company with the given parameters and returns the created company's internal DTO.
	CreateCompany(body userdto.CreateCompanyParams) (userdto.CompanyInternalDTO, error)

	// GetAllCompanies retrieves all companies and returns a slice of company internal DTOs.
	GetAllCompanies() ([]userdto.CompanyInternalDTO, error)

	// GetCompanyByName retrieves a company by its name and returns the company's internal DTO.
	GetCompanyByName(name string) (userdto.CompanyInternalDTO, error)

	// GetCompanyByID retrieves a company by its ID and returns the company's internal DTO.
	GetCompanyByID(id string) (userdto.CompanyInternalDTO, error)
}

// CompanyService method
func (s *service) CompanyService() CompanyService {
	return serviceProvider(s)
}

type PaymentService interface {
	GetBankList() ([]Bank, error)
}

// Payment service
func (s *service) NewPaymentService() PaymentService {
	return serviceProvider(s)
}
