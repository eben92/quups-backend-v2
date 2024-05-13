package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	repository "quups-backend/internal/database/repository"
	"quups-backend/internal/util"
)

type Service struct {
	Repo *repository.Queries
}

func (s *Service) Signin(w http.ResponseWriter, r *http.Request) {

	response := util.Response{}
	var res []byte

	u, err := s.Repo.GetUsers(r.Context())

	if err != nil {
		fmt.Print(err.Error())

		res, _ = response.Builder(w, r, &util.ApiResponseParams{
			Results:    err.Error(),
			StatusCode: http.StatusBadRequest,
			Message:    util.String("Error getting users. Please try again"),
		})

		_, _ = w.Write(res)
		return
	}

	result := []userDTO{}

	for i := 0; i < len(u); i++ {
		u := mapToUserDTO(u[i])
		result = append(result, *u)
	}

	// w.WriteHeader(http.StatusCreated)
	res, _ = response.Builder(w, r, &util.ApiResponseParams{
		Results:    result,
		StatusCode: http.StatusOK,
		Message:    util.String("users retrieved successfully"),
	})

	_, _ = w.Write(res)
}

type userBody struct {
	Email    *string
	Name     *string
	Msisdn   string
	Password string
}

func (s *Service) Signup(w http.ResponseWriter, r *http.Request) {
	var body *userBody
	response := util.Response{}

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		res, _ := response.Builder(w, r, &util.ApiResponseParams{
			Results:    nil,
			StatusCode: http.StatusBadRequest,
			Message:    util.String(err.Error()),
		})

		_, _ = w.Write(res)
		return
	}

	// check if email is nil
	if body.Email == nil {
		w.WriteHeader(http.StatusBadRequest)

		res, _ := response.Builder(w, r, &util.ApiResponseParams{
			Results:    nil,
			StatusCode: http.StatusBadRequest,
			Message:    util.String("Email field is required"),
		})

		_, _ = w.Write(res)
		return
	}

	/*todo:
	- check to see if email and msisdn already exists and throw error  if it does
	- create user and generate jwt signed token
	- send the signed token in both the request body and append it to the browser cookie
	-
	*/

	//  save user in db
	u, err := s.createUser(r.Context(), body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		res, _ := response.Builder(w, r, &util.ApiResponseParams{
			Results:    nil,
			StatusCode: http.StatusBadRequest,
			Message:    util.String(err.Error()),
		})

		_, _ = w.Write(res)
		return
	}

	new_u := mapToUserDTO(u)

	res, _ := response.Builder(w, r, &util.ApiResponseParams{
		Results:    new_u,
		StatusCode: http.StatusBadRequest,
		Message:    util.String("user created successfully"),
	})

	_, _ = w.Write(res)
}

func (s *Service) createUser(ctx context.Context, body *userBody) (repository.User, error) {

	n := *body.Name


	u, err := s.Repo.CreateUser(ctx, repository.CreateUserParams{
		Email: *body.Email,
		Name: sql.NullString{
			String: n,
			Valid:  strconv.ParseBool(n); err != nil {
				return false
			},
		},
	})

	if err != nil {

		return repository.User{}, err
	}

	return u, nil

}
