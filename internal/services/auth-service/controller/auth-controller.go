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

	"github.com/go-chi/chi/v5"
)

type AuthController interface {
	Signin(w http.ResponseWriter, r *http.Request)
	AccountSignin(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
	Signout(w http.ResponseWriter, r *http.Request)
	GetUserCompany(w http.ResponseWriter, r *http.Request)
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
	aservice := authservice.NewAuthService(r.Context(), s.db)
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
	setCookie(w, Cookie{
		Value: access_token,
	})

	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Results:    user,
		Message:    success,
	})
}

// POST: /user/account
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
	setCookie(w, Cookie{

		Value: access_token,
	})

	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Results:    success,
		Message:    success,
	})
}

// POST: /auth/signup
func (s *Controller) Signup(w http.ResponseWriter, r *http.Request) {
	body := userdto.CreateUserParams{}

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
	aservice := authservice.NewAuthService(r.Context(), s.db)
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

		// todo: remove token session from redis
		setCookie(w, Cookie{
			Value:   "",
			Name:    local_jwt.COOKIE_NAME_COMPANY,
			Expires: time.Now().Add(-time.Hour),
		})
		w.WriteHeader(http.StatusNoContent)
		return

	}

	setCookie(w, Cookie{
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})
	w.WriteHeader(http.StatusNoContent)
}

func (c *Controller) GetUserCompany(w http.ResponseWriter, r *http.Request) {
	response := apiutils.New(w, r)
	usrv := userservice.NewUserService(r.Context(), c.db)

	companyID := chi.URLParam(r, "id")

	t, err := usrv.GetUserTeam(companyID)

	if err != nil {

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusOK,
			Message:    err.Error(),
			Results:    nil,
		})

		return
	}

	// todo: save this token as a session in redis
	token, err := local_jwt.GenereteJWT(local_jwt.AuthContext{
		Sub:       t.Company.ID,
		Name:      t.Company.Name,
		CompanyID: t.Company.ID,
	})

	if err != nil {
		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusOK,
			Message:    err.Error(),
			Results:    nil,
		})

		return
	}

	t.AccessToken = string(token)

	setCookie(w, Cookie{
		Name:  local_jwt.COOKIE_NAME_COMPANY,
		Value: string(token),
	})

	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success",
		Results:    t,
	})
}

type Cookie struct {
	Name    local_jwt.COOKIE_NAME
	Value   string
	Expires time.Time
}

func setCookie(w http.ResponseWriter, c Cookie) {
	exp := c.Expires

	if exp.IsZero() {
		exp = time.Now().Add(time.Hour * 24 * 30)
	}

	name := local_jwt.COOKIE_NAME_USER

	if c.Name != "" {
		name = c.Name
	}

	cookie := &http.Cookie{
		Name:     string(name),
		Value:    c.Value,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		HttpOnly: false,
		Expires:  exp,
		// Domain:   "*",
	}

	http.SetCookie(w, cookie)
}
