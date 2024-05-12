package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	repository "quups-backend/internal/database/repository"
)

type Service struct {
	Repo *repository.Queries
}

func (s *Service) Signin(w http.ResponseWriter, r *http.Request) {

	res, err := s.Repo.FindMany(r.Context())

	if err != nil {
		fmt.Print(err.Error())

		_, _ = w.Write([]byte(err.Error()))
		return
	}

	fmt.Printf("login successfully")

	var response struct {
		Results []repository.User `json:"results"`
	}
	response.Results = res

	data, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(data)
}

type userDTO struct {
	Email *string
}

type ApiError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Result     any    `json:"result"`
}

func (s *Service) Signup(w http.ResponseWriter, r *http.Request) {
	var body *userDTO

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
	u, err := s.Repo.Create(r.Context(), *body.Email)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	res, _ := json.Marshal(u)

	fmt.Println("signup successfully")
	_, _ = w.Write(res)
}
