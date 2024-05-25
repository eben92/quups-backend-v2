package usercontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	userdto "quups-backend/internal/services/user-service/dto"
	userservice "quups-backend/internal/services/user-service/service"
	apiutils "quups-backend/internal/utils/api"
	local_jwt "quups-backend/internal/utils/jwt"
)

// POST: /companies
func (c *Controller) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var body *userdto.CreateCompanyParams
	uservice := userservice.New(r.Context(), c.repo)
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

	newc, err := uservice.CreateCompany(body)
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
func (c *Controller) GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	response := apiutils.New(w, r)
	uservice := userservice.New(r.Context(), c.repo)

	companies, err := uservice.GetAllCompanies()
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
func (c *Controller) GetCompanyByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	name := r.URL.Query().Get("name")

	log.Printf("name :[%s]", name)
	uservice := userservice.New(r.Context(), c.repo)
	response := apiutils.New(w, r)

	co, err := uservice.GetCompanyByID(id)
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
func (c *Controller) GetCompanyByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	uservice := userservice.New(r.Context(), c.repo)
	response := apiutils.New(w, r)

	co, err := uservice.GetCompanyByName(name)
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

func (c *Controller) GetUserTeams(w http.ResponseWriter, r *http.Request) {
	response := apiutils.New(w, r)

	claims := local_jwt.GetAuthContext(r.Context())

	usrv := userservice.New(r.Context(), c.repo)

	t, err := usrv.GetUserTeams(claims.Sub)
	if err != nil {

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusOK,
			Message:    err.Error(),
			Results:    nil,
		})

		return
	}
	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success",
		Results:    t,
	})
}
