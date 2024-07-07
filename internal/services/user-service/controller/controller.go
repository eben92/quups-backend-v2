package usercontroller

import (
	"net/http"

	"quups-backend/internal/database"
)

//	type Controller interface {
//		NewCompanyController() Company
//	}
type controller struct {
	db database.Service
}

// fun New(db *database.Service) Controller {
// 	return &controller{
// 		db: db,
// 	}
// }

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
	GetUserTeams(w http.ResponseWriter, r *http.Request)
}

func NewUserController(db database.Service) userController {
	return &controller{
		db: db,
	}
}

type paymentController interface {
	GetBankList(w http.ResponseWriter, r *http.Request)
}

func NewPaymentController(db database.Service) paymentController {
	return &controller{
		db: db,
	}
}
