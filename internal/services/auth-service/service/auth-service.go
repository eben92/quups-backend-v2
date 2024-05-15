package authservice

import (
	"context"
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
	ctx  context.Context
}

func New(ctx context.Context, r *repository.Queries) *Service {
	return &Service{
		repo: r,
		ctx:  ctx,
	}
}

func (s *Service) SigninHandler(body *authdto.SignInRequestDTO) (*authdto.ResponseUserDTO, error) {
	uService := userservice.New(s.ctx, s.repo)

	u, err := uService.FindByMsisdn(body.Msisdn)

	if err != nil {
		return nil, fmt.Errorf("incorrect phone number or password")
	}

	user := mapToUserDTO(u)

	return user, nil
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

func mapToUserDTO(user *userdto.UserInternalDTO) *authdto.ResponseUserDTO {

	dto := &authdto.ResponseUserDTO{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Msisdn:   user.Msisdn,
		ImageUrl: user.ImageUrl,
		Gender:   user.Gender,
	}

	return dto

}
