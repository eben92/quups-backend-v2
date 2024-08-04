package usercontroller

import (
	"net/http"

	"quups-backend/internal/database"
)

type controller struct {
	db database.Service
}

type companyController interface {
	CreateCompany(w http.ResponseWriter, r *http.Request)
	GetCompanyByID(w http.ResponseWriter, r *http.Request)
	GetCompanyByName(w http.ResponseWriter, r *http.Request)
	GetCompanyNameAvailability(w http.ResponseWriter, r *http.Request)
	GetAllCompanies(w http.ResponseWriter, r *http.Request)
}

func NewCompanyController(db database.Service) companyController {
	return &controller{
		db: db,
	}
}

type userController interface {
	GetUserCompanies(w http.ResponseWriter, r *http.Request)
	GetUserCompany(w http.ResponseWriter, r *http.Request)
}

func NewUserController(db database.Service) userController {
	return &controller{
		db: db,
	}
}
