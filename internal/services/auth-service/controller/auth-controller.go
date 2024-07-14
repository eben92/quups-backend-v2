package authcontroller

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"quups-backend/internal/database"
	authdto "quups-backend/internal/services/auth-service/dto"
	authservice "quups-backend/internal/services/auth-service/service"
	userdto "quups-backend/internal/services/user-service/dto"
	userservice "quups-backend/internal/services/user-service/service"
	apiutils "quups-backend/internal/utils/api"
	local_jwt "quups-backend/internal/utils/jwt"
)

type AuthController interface {
	Signin(w http.ResponseWriter, r *http.Request)
	AccountSignin(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
	Signout(w http.ResponseWriter, r *http.Request)
}

type Controller struct {
	db database.Service
}

const (
	invalidRequest = "Invalid request body."
	success        = "success"
)

func New(db database.Service) AuthController {
	return &Controller{
		db: db,
	}
}

// POST: /auth/signin
func (s *Controller) Signin(w http.ResponseWriter, r *http.Request) {
	var body authdto.SignInRequestDTO
	aservice := authservice.NewAuthService(r.Context(), s.db)
	response := apiutils.New(w, r)

	err := json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	if err != nil {
		slog.Error("error decoding signin request body", "Error", err)

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    invalidRequest,
		})

		return
	}

	user, err := aservice.Signin(body)
	if err != nil {
		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    err.Error(),
		})

		return
	}

	access_token := user.AccessToken
	setCookie(w, access_token)

	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Results:    user,
		Message:    success,
	})
}

// POST: /auth/account
func (s *Controller) AccountSignin(w http.ResponseWriter, r *http.Request) {
	var body authdto.AccountSigninDTO
	aservice := authservice.NewAuthService(r.Context(), s.db)
	response := apiutils.New(w, r)

	err := json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	if err != nil {
		slog.Error("error decoding signin request body", "Error", err)

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    invalidRequest,
		})

		return
	}

	user, err := aservice.AccountSignin(body)

	if err != nil {
		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusUnauthorized,
			Results:    nil,
			Message:    err.Error(),
		})

		return
	}

	access_token := user.AccessToken
	setCookie(w, access_token)

	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Results:    success,
		Message:    success,
	})
}

// POST: /auth/signup
func (s *Controller) Signup(w http.ResponseWriter, r *http.Request) {
	body := userdto.CreateUserParams{}
	aservice := authservice.NewAuthService(r.Context(), s.db)
	response := apiutils.New(w, r)

	err := json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	if err != nil {
		slog.Error("error decoding signin request body", "Error", err)

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    invalidRequest,
		})

		return
	}

	if err := userservice.ValidateCreateUserQ(body); err != nil {
		slog.Error("error decoding signin request body", "Error", err)

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    invalidRequest,
		})

		return
	}

	user, err := aservice.Signup(body)

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
		Results:    user, // TODO: shoudld we take this out?
		Message:    success,
	})
}

// POST: /auth/signout
func (s *Controller) Signout(w http.ResponseWriter, r *http.Request) {
	querytype := r.URL.Query().Get("type")

	if strings.ToLower(querytype) == "soft" {
		aservice := authservice.NewAuthService(r.Context(), s.db)

		tstring, err := aservice.SoftSignout()

		if err != nil {
			response := apiutils.New(w, r)
			response.WrapInApiResponse(&apiutils.ApiResponseParams{
				StatusCode: http.StatusForbidden,
				Results:    nil,
				Message:    "error signing out",
			})

			return
		}

		setCookie(w, tstring)
		response := apiutils.New(w, r)
		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusOK,
			Results:    nil,
			Message:    success,
		})

		return
	}

	cookie := &http.Cookie{
		Name:    local_jwt.COOKIE_NAME,
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	}

	http.SetCookie(w, cookie)

	response := apiutils.New(w, r)
	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Results:    nil,
		Message:    success,
	})

}

func setCookie(w http.ResponseWriter, t string) {
	cookie := &http.Cookie{
		Name:     local_jwt.COOKIE_NAME,
		Value:    t,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		HttpOnly: false,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		// Domain:   "*",
	}

	http.SetCookie(w, cookie)
}
