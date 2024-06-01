package authcontroller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"quups-backend/internal/database"
	authdto "quups-backend/internal/services/auth-service/dto"
	authservice "quups-backend/internal/services/auth-service/service"
	userdto "quups-backend/internal/services/user-service/dto"
	apiutils "quups-backend/internal/utils/api"
	local_jwt "quups-backend/internal/utils/jwt"
)

type Controller struct {
	db database.Service
}

const (
	invalidRequest = "Invalid request body."
	success        = "success"
)

func New(db database.Service) *Controller {
	return &Controller{
		db: db,
	}
}

// POST: /auth/signin
func (s *Controller) Signin(w http.ResponseWriter, r *http.Request) {
	var body *authdto.SignInRequestDTO
	aservice := authservice.New(r.Context(), s.db)
	response := apiutils.New(w, r)

	err := json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	if err != nil {
		log.Printf("error decoding signin request body")

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    invalidRequest,
		})

		return
	}

	user, err := aservice.SigninHandler(body)
	if err != nil {
		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    err.Error(),
		})

		return
	}

	access_token := user.AccessToken
	setCookie(w, *access_token)

	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Results:    &user,
		Message:    success,
	})
}

// POST: /auth/signup
func (s *Controller) Signup(w http.ResponseWriter, r *http.Request) {
	var body *userdto.CreateUserParams
	aservice := authservice.New(r.Context(), s.db)
	response := apiutils.New(w, r)

	err := json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	if err != nil {
		log.Printf("error decoding signin request body")

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    invalidRequest,
		})

		return
	}

	user, err := aservice.SignupHandler(body)
	if err != nil {
		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    err.Error(),
		})

		return
	}

	// TODO:
	// add OTP and redirect user to confirm their phone number

	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusCreated,
		Results:    &user, // TODO: shoudld we take this out?
		Message:    success,
	})
}

func setCookie(w http.ResponseWriter, t string) {
	cookie := &http.Cookie{
		Name:     local_jwt.COOKIE_NAME,
		Value:    t,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		// Domain:   "*",
	}

	http.SetCookie(w, cookie)
}
