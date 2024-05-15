package authcontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"quups-backend/internal/database/repository"
	authdto "quups-backend/internal/services/auth-service/dto"
	authservice "quups-backend/internal/services/auth-service/service"
	userdto "quups-backend/internal/services/user-service/dto"
	"quups-backend/internal/utils"
)

type Controller struct {
	repo *repository.Queries
}

const (
	invalidRequest = "Invalid request body."
	success        = "success"
)

func New(r *repository.Queries) *Controller {
	return &Controller{
		repo: r,
	}
}

// POST: /auth/signin
func (s *Controller) Signin(w http.ResponseWriter, r *http.Request) {
	var body *authdto.SignInRequestDTO
	aservice := authservice.New(r.Context(), s.repo)
	response := utils.New(w, r)

	err := json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	if err != nil {
		log.Printf("error decoding signin request body")

		res, _ := response.WrapInApiResponse(&utils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    invalidRequest,
		})

		_, _ = w.Write(res)
		return
	}

	user, err := aservice.SigninHandler(body)

	if err != nil {
		res, _ := response.WrapInApiResponse(&utils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    err.Error(),
		})

		_, _ = w.Write(res)
		return
	}

	res, _ := response.WrapInApiResponse(&utils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Results:    &user,
		Message:    success,
	})

	_, _ = w.Write(res)

}

// POST: /auth/signup
func (s *Controller) Signup(w http.ResponseWriter, r *http.Request) {
	var body *userdto.CreateUserParams
	aservice := authservice.New(r.Context(), s.repo)
	response := utils.New(w, r)

	err := json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	if err != nil {
		log.Printf("error decoding signin request body")

		res, _ := response.WrapInApiResponse(&utils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    invalidRequest,
		})

		_, _ = w.Write(res)
		return
	}

	user, err := aservice.SignupHandler(body)

	if err != nil {
		res, _ := response.WrapInApiResponse(&utils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    err.Error(),
		})

		_, _ = w.Write(res)
		return
	}

	res, _ := response.WrapInApiResponse(&utils.ApiResponseParams{
		StatusCode: http.StatusCreated,
		Results:    &user,
		Message:    success,
	})

	_, _ = w.Write(res)

}
