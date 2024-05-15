package authcontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"quups-backend/internal/database/repository"
	authdto "quups-backend/internal/services/auth-service/dto"
	authservice "quups-backend/internal/services/auth-service/service"
	"quups-backend/internal/utils"
)

type Controller struct {
	repo *repository.Queries
}

func New(r *repository.Queries) *Controller {
	return &Controller{
		repo: r,
	}
}

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
			Message:    "Invalid request body.",
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
		StatusCode: http.StatusBadRequest,
		Results:    &user,
		Message:    "success",
	})

	_, _ = w.Write(res)

}

func (s *Controller) Signup(w http.ResponseWriter, r *http.Request) {

}
