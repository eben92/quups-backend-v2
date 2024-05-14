package authservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"quups-backend/internal/database/repository"
	authdto "quups-backend/internal/services/auth-service/dto"
	userdto "quups-backend/internal/services/user-service/dto"
	userservice "quups-backend/internal/services/user-service/service"
	"quups-backend/internal/utils"
)

const (
	emailErr    = "Email is required"
	nameErr     = "Name is required"
	enErr       = "Email and name is required"
	msisdnTaken = "Phone number already in use."
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

	result := []authdto.UserDTO{}

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
	var user *userdto.UserInternalDTO

	util := utils.New(w, r)
	uService := userservice.New(r.Context(), s.repo)

	err := json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		res, _ := util.WrapInApiResponse(&utils.ApiResponseParams{
			Results:    nil,
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})

		_, _ = w.Write(res)
		return
	}

	// check to see if email or msisdn already exists and throw error  if it does

	// if body.Email == nil {
	// 	log.Println(emailErr)

	// 	w.WriteHeader(http.StatusBadRequest)
	// 	res, _ := util.WrapInApiResponse(&utils.ApiResponseParams{
	// 		Results:    nil,
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    emailErr,
	// 	})

	// 	_, _ = w.Write(res)
	// 	return

	// }

	//create user and generate jwt signed token
	// send the signed token in both the request body and append it to the browser cookie

	//  save user in db

	user, err = uService.Create(body)

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
		Results:    user,
		StatusCode: http.StatusOK,
		Message:    "user created successfully",
	})

	_, _ = w.Write(res)
}

func mapToUserDTO(user repository.User) *authdto.UserDTO {

	dto := &authdto.UserDTO{
		ID:    user.ID,
		Email: user.Email,
	}

	if user.Name.Valid {
		dto.Name = &user.Name.String
	}

	if user.Msisdn.Valid {
		dto.Msisdn = &user.Msisdn.String
	}

	if user.ImageUrl.Valid {
		dto.ImageUrl = &user.ImageUrl.String
	}

	if user.Gender.Valid {
		dto.Gender = &user.Gender.String
	}

	return dto

}
