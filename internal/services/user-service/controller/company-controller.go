package usercontroller

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	userdto "quups-backend/internal/services/user-service/dto"
	userservice "quups-backend/internal/services/user-service/service"
	"quups-backend/internal/utils"
	apiutils "quups-backend/internal/utils/api"
)

// POST: /companies
func (c *controller) CreateCompany(w http.ResponseWriter, r *http.Request) {
	body := userdto.CreateCompanyParams{}
	cmpsrv := userservice.NewCompanyService(r.Context(), c.db)
	response := apiutils.New(w, r)

	err := json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	if err != nil {
		slog.Error("error decoding create company request body", "Error", err)

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    err.Error(),
		})
		return
	}

	if err := userservice.ValidateCreateCompanyQParams(body); err != nil {
		slog.Error("error validating create company request body", "Error", err)

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    err.Error(),
		})
		return
	}

	result, err := cmpsrv.CreateCompany(body)
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
		Results:    result,
		Message:    "success",
	})
}

// GET: /companies
func (c *controller) GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	response := apiutils.New(w, r)
	cmpsrv := userservice.NewCompanyService(r.Context(), c.db)

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

	cmpsrv := userservice.NewCompanyService(r.Context(), c.db)
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

	cmpsrv := userservice.NewCompanyService(r.Context(), c.db)
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

// GET: /companies/name/exists?name=""
func (c *controller) GetCompanyNameAvailability(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	response := apiutils.New(w, r)

	type Response struct {
		Message      string `json:"message"`
		Availability string `json:"availability"`
	}

	message := name + " already in use. please choose another"
	availability := "NOT_AVAILABLE"

	if name == "" {
		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusNotFound,
			Message:    "name query params is required",
			Results:    nil,
		})
		return
	}

	n, isvalid := utils.IsValidCompanyName(name)

	if !isvalid {
		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusOK,
			Message:    "Invalid store name. Please choose another",
			Results: Response{
				Message:      "Company name cannot contain space, special characters or accented letters.",
				Availability: availability,
			},
		})
		return
	}

	cmpsrv := userservice.NewCompanyService(r.Context(), c.db)

	co, err := cmpsrv.GetCompanyByName(n)

	if co.ID == "" {
		message = n + " is available"
		availability = "AVAILABLE"
	}

	res := Response{
		Message:      message,
		Availability: availability,
	}

	if err != nil {
		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusOK,
			Message:    err.Error(),
			Results:    res,
		})

		return
	}

	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success",
		Results:    res,
	})
}
