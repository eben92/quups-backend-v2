package authservice

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	repository "quups-backend/internal/database/repository"
	userservice "quups-backend/internal/services/user-service"
	"quups-backend/internal/util"
)

type Service struct {
	repo *repository.Queries
}

func New(r *repository.Queries) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Signin(w http.ResponseWriter, r *http.Request) {

	response := util.Response{}
	var res []byte

	u, err := s.repo.GetUsers(r.Context())

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

	fmt.Println(result)

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
	Msisdn   *string
	Password string
}

func (s *Service) Signup(w http.ResponseWriter, r *http.Request) {
	var body *userBody
	response := util.Response{}
	userService := userservice.New(s.repo)

	userService.GetUserByEmail()

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
	if body.Email == nil || body.Name == nil || body.Msisdn == nil {
		w.WriteHeader(http.StatusBadRequest)

		res, _ := response.Builder(w, r, &util.ApiResponseParams{
			Results:    nil,
			StatusCode: http.StatusBadRequest,
			Message:    util.String("Email and Name is required"),
		})

		_, _ = w.Write(res)
		return
	}

	// check to see if email or msisdn msisdn already exists and throw error  if it does
	g, err := s.repo.GetUserByEmail(r.Context(), *body.Email)

	fmt.Println(g)

	cUser := userDTO{}

	//*mapToUserDTO(userInDB)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		res, _ := response.Builder(w, r, &util.ApiResponseParams{
			Results:    nil,
			StatusCode: http.StatusBadRequest,
			Message:    util.String(err.Error()),
		})

		_, _ = w.Write(res)
	}

	// if cUser != nil {
	// 	w.WriteHeader(http.StatusBadRequest)

	// 	res, _ := response.Builder(w, r, &util.ApiResponseParams{
	// 		Results:    nil,
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    util.String(err.Error()),
	// 	})

	// 	_, _ = w.Write(res)
	// }

	//create user and generate jwt signed token
	// send the signed token in both the request body and append it to the browser cookie

	//  save user in db
	// u, err := s.createUser(r.Context(), body)

	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)

	// 	res, _ := response.Builder(w, r, &util.ApiResponseParams{
	// 		Results:    nil,
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    util.String(err.Error()),
	// 	})

	// 	_, _ = w.Write(res)
	// 	return
	// }

	// new_u := mapToUserDTO(u)

	res, _ := response.Builder(w, r, &util.ApiResponseParams{
		Results:    cUser,
		StatusCode: http.StatusBadRequest,
		Message:    util.String("user created successfully"),
	})

	_, _ = w.Write(res)
}

func (s *Service) createUser(ctx context.Context, body *userBody) (repository.User, error) {

	n := *body.Name

	u, err := s.repo.CreateUser(ctx, repository.CreateUserParams{
		Email: *body.Email,
		Name: sql.NullString{
			String: n,
			Valid:  true,
		},
	})

	if err != nil {

		return repository.User{}, err
	}

	return u, nil

}
