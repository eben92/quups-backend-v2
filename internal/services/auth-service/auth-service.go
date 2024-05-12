package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	repository "quups-backend/internal/database/repository"
)

type Service struct {
	Repo *repository.Queries
}

func (s *Service) Signin(w http.ResponseWriter, r *http.Request) {

	res, err := s.Repo.GetUsers(r.Context())

	if err != nil {
		fmt.Print(err.Error())

		_, _ = w.Write([]byte(err.Error()))
		return
	}

	result := []userDTO{}
	for i := 0; i < len(res); i++ {
		u := mapToUserDTO(res[i])
		result = append(result, *u)
	}

	var response struct {
		Results []userDTO `json:"results"`
	}
	response.Results = result

	data, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(data)
}

type userBody struct {
	Email *string
}

type ApiError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Result     any    `json:"result"`
}

func (s *Service) Signup(w http.ResponseWriter, r *http.Request) {
	var body *userBody

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, _ = w.Write([]byte(err.Error()))
		return
	}

	// check if email is nil
	if body.Email == nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(&ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    "Email field is required",
			Result:     nil,
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
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	new_u := mapToUserDTO(u)
	res, _ := json.Marshal(new_u)

	_, _ = w.Write(res)
}

func (s *Service) createUser(ctx context.Context, body *userBody) (repository.User, error) {

	u, err := s.Repo.CreateUser(ctx, repository.CreateUserParams{
		Email: *body.Email,
	})

	if err != nil {

		return repository.User{}, err
	}

	return u, nil

}
