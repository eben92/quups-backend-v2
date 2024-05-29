package usercontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	userdto "quups-backend/internal/services/user-service/dto"
	userservice "quups-backend/internal/services/user-service/service"
	apiutils "quups-backend/internal/utils/api"
)

// POST: /companies
func (c *controller) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var body *userdto.CreateCompanyParams
	cmpsrv := userservice.New(r.Context(), c.db).CompanyService()
	response := apiutils.New(w, r)

	err := json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	if err != nil {
		log.Printf("error decoding create company request body")

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    err.Error(),
		})
		return
	}

	newc, err := cmpsrv.CreateCompany(body)
	if err != nil {

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    err.Error(),
		})
		return
	}

	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusCreated,
		Results:    &newc,
		Message:    "success",
	})
}

// GET: /companies
func (c *controller) GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	response := apiutils.New(w, r)
	cmpsrv := userservice.New(r.Context(), c.db).CompanyService()

	companies, err := cmpsrv.GetAllCompanies()
	if err != nil {

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    err.Error(),
		})
		return

	}

	log.Println(companies)

	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Results:    companies,
		Message:    "success",
	})
}

// GET: /companies/{id}
func (c *controller) GetCompanyByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	name := r.URL.Query().Get("name")

	log.Printf("name :[%s]", name)
	cmpsrv := userservice.New(r.Context(), c.db).CompanyService()
	response := apiutils.New(w, r)

	co, err := cmpsrv.GetCompanyByID(id)
	if err != nil {
		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Results:    nil,
		})

		return
	}

	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success",
		Results:    co,
	})
}

// GET: /companies/name/{name}
func (c *controller) GetCompanyByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	cmpsrv := userservice.New(r.Context(), c.db).CompanyService()
	response := apiutils.New(w, r)

	co, err := cmpsrv.GetCompanyByName(name)
	if err != nil {
		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Results:    nil,
		})

		return
	}

	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success",
		Results:    co,
	})
}
