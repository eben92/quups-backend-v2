package authservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"quups-backend/internal/database/repository"
	userdto "quups-backend/internal/services/user-service/dto"
	userservice "quups-backend/internal/services/user-service/service"
	"quups-backend/internal/utils"
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
	util := utils.New(w, r)

	var res []byte

	u, err := s.repo.GetUsers(r.Context())

	if err != nil {
		fmt.Print(err.Error())

		res, _ = util.WrapInApiResponse(&utils.ApiResponseParams{
			Results:    err.Error(),
			StatusCode: http.StatusBadRequest,
			Message:    "Error getting users. Please try again",
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
	res, _ = util.WrapInApiResponse(&utils.ApiResponseParams{
		Results:    result,
		StatusCode: http.StatusOK,
		Message:    "users retrieved successfully",
	})

	_, _ = w.Write(res)
}

func (s *Service) Signup(w http.ResponseWriter, r *http.Request) {
	var body *userdto.CreateUserParams
	util := utils.New(w, r)
	uService := userservice.New(r.Context(), s.repo)

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		fmt.Println("error herrrrrrrrrrrrrrrrrrrrrrrrrr")
		w.WriteHeader(http.StatusBadRequest)

		res, _ := util.WrapInApiResponse(&utils.ApiResponseParams{
			Results:    nil,
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})

		_, _ = w.Write(res)
		return
	}

	// check if email is nil
	// if body.Email == nil || body.Name == nil || body.Msisdn == nil {

	// 	res, _ := util.WrapInApiResponse(&utils.ApiResponseParams{
	// 		Results:    nil,
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    "Email and Name is required",
	// 	})

	// 	_, _ = w.Write(res)
	// 	return
	// }

	// check to see if email or msisdn msisdn already exists and throw error  if it does
	// cUser, err := uService.FindByEmail(*body.Email)

	// if err != nil {
	// 	res, _ := util.WrapInApiResponse(&utils.ApiResponseParams{
	// 		Results:    nil,
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    err.Error(),
	// 	})

	// 	_, _ = w.Write(res)
	// 	return
	// }

	//create user and generate jwt signed token
	// send the signed token in both the request body and append it to the browser cookie

	//  save user in db

	cUser, err := uService.Create(body)

	if err != nil {
		res, _ := util.WrapInApiResponse(&utils.ApiResponseParams{
			Results:    nil,
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})

		_, _ = w.Write(res)
		return
	}

	res, _ := util.WrapInApiResponse(&utils.ApiResponseParams{
		Results:    cUser,
		StatusCode: http.StatusOK,
		Message:    "user created successfully",
	})

	_, _ = w.Write(res)
}
