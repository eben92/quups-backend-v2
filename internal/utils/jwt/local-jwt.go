package local_jwt

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	apiutils "quups-backend/internal/utils/api"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type COOKIE_NAME string

const (
	COOKIE_NAME_USER    COOKIE_NAME = "qp-session"
	COOKIE_NAME_COMPANY COOKIE_NAME = "qp-client"
)

var (
	JWT_SECRET      = os.Getenv("JWT_SECRET")
	AUTH_CTX_KEY    = &ContextKey{"authcontext"}
	COMPANY_CTX_KEY = &ContextKey{"companycontext"}
)

type AuthContext struct {
	// Sub is the subject of the token
	// This is used to determine the user the token belongs to
	Sub string

	// Issuer is the issuer of the token
	// This is used to determine the source of the token
	Issuer string

	// Name is the name of the user
	Name string

	// CompanyID is the company id the current user has logged in to
	// This is used to determine the company the user is currently working with
	// in JWT, this is the 'client_id'
	CompanyID string
}

type ContextKey struct {
	name string
}

type tokenMgt struct {
	UserToken    string
	CompanyToken string
}

var (
	ErrUnauthorized = errors.New("token is unauthorized")
	ErrExpired      = errors.New("token is expired")
	ErrNBFInvalid   = errors.New("token nbf validation failed")
	ErrIATInvalid   = errors.New("token iat validation failed")
	ErrNoTokenFound = errors.New("no token found")
	ErrAlgoInvalid  = errors.New("algorithm mismatch")
)

// UserMiddleware is a middleware that handles user jwt authentications
func UserMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			var token tokenMgt

			findtokens := []func(*http.Request) tokenMgt{GetTokenFromHeader, GetTokenFromCookie}

			w.Header().Set("Content-Type", "application/json")

			for _, fn := range findtokens {
				token = fn(r)

				if token.UserToken != "" {
					break
				}

			}

			if token.UserToken == "" {

				response := apiutils.New(w, r)

				response.WrapInApiResponse(&apiutils.ApiResponseParams{
					StatusCode: http.StatusUnauthorized,
					Message:    http.StatusText(http.StatusUnauthorized),
					Results:    nil,
				})

				return

			}

			c, err := ParseToken(token.UserToken)
			if err != nil {
				response := apiutils.New(w, r)

				response.WrapInApiResponse(&apiutils.ApiResponseParams{
					StatusCode: http.StatusUnauthorized,
					Message:    err.Error(),
					Results:    nil,
				})

				return
			}

			ctx, err := newContext(r.Context(), AUTH_CTX_KEY, c)
			if err != nil {

				// http.Error(w, err.Error(), http.StatusUnauthorized)

				response := apiutils.New(w, r)

				response.WrapInApiResponse(&apiutils.ApiResponseParams{
					StatusCode: http.StatusUnauthorized,
					Message:    err.Error(),
					Results:    nil,
				})

				return

			}

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(hfn)
	}
}

// CompanyMiddleware is a middleware that handles company jwt authentications
func CompanyMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			var token tokenMgt

			findtokens := []func(*http.Request) tokenMgt{GetTokenFromHeader, GetTokenFromCookie}

			for _, fn := range findtokens {
				token = fn(r)

				if token.CompanyToken != "" {
					break
				}

			}

			if token.CompanyToken == "" {

				response := apiutils.New(w, r)

				response.WrapInApiResponse(&apiutils.ApiResponseParams{
					StatusCode: http.StatusUnauthorized,
					Message:    http.StatusText(http.StatusUnauthorized),
					Results:    nil,
				})

				return

			}

			c, err := ParseToken(token.CompanyToken)
			if err != nil {
				response := apiutils.New(w, r)

				response.WrapInApiResponse(&apiutils.ApiResponseParams{
					StatusCode: http.StatusUnauthorized,
					Message:    err.Error(),
					Results:    nil,
				})

				return
			}

			ctx, err := newContext(r.Context(), COMPANY_CTX_KEY, c)
			if err != nil {

				response := apiutils.New(w, r)

				response.WrapInApiResponse(&apiutils.ApiResponseParams{
					StatusCode: http.StatusUnauthorized,
					Message:    err.Error(),
					Results:    nil,
				})

				return

			}

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(hfn)
	}
}

func GetTokenFromHeader(r *http.Request) tokenMgt {
	bearer := r.Header.Get("Authorization")
	t := tokenMgt{}

	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		slog.Info("got token from request head")

		companyToken := r.Header.Get(string(COOKIE_NAME_COMPANY))

		if companyToken != "" {
			slog.Warn("company token found in header")
		}

		return tokenMgt{bearer[7:], companyToken}

	}

	return t
}

func GetTokenFromCookie(r *http.Request) tokenMgt {
	t := tokenMgt{}
	userCookie, err := r.Cookie(string(COOKIE_NAME_USER))

	if err != nil {
		return t
	}

	t.UserToken = userCookie.Value

	companyCookie, err := r.Cookie(string(COOKIE_NAME_COMPANY))

	if err != nil {

		slog.Warn("no company cookie found")

		return t

	}

	t.CompanyToken = companyCookie.Value

	slog.Info("user and company token found in cookie")

	return t
}

/*
Generetes a signed token and return as byte or nil.
Convert to string before sending to client
*/
func GenereteJWT(data AuthContext) ([]byte, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"service_id": "quups-backend",
		"sub":        data.Sub,
		"client_id":  data.CompanyID,
		"name":       data.Name,
		"issuer":     "WEB",
		"exp":        time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		slog.Error("Error signing jwt", "Error", err)

		return nil, fmt.Errorf("Something went wrong. Please try again. #2")
	}

	return []byte(tokenString), nil
}

// ParseToken is use to verify jwt tokens
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, alg := token.Method.(*jwt.SigningMethodHMAC); !alg {

			slog.Warn("invalid token alg", "ParseToken", token.Header["alg"])

			return nil, ErrAlgoInvalid
		}

		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		slog.Error("error parsing token [%s]", "Error", err)

		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, ErrNBFInvalid
	}

	return claims, nil
}

func newContext(ctx context.Context, key *ContextKey, claims jwt.MapClaims) (context.Context, error) {
	sub := claims["sub"]

	if sub == "" {
		return nil, ErrUnauthorized
	}

	ctx = context.WithValue(ctx, key, claims)
	return ctx, nil
}

// GetAuthContext returns decoded jwt data
func GetAuthContext(ctx context.Context, key *ContextKey) (AuthContext, error) {
	claims, ok := ctx.Value(key).(jwt.MapClaims)

	if !ok {
		slog.Error("GetAuthContext - no claims found", "Error", ErrNoTokenFound)

		return AuthContext{}, ErrNoTokenFound
	}

	var companyID string
	var sub string
	var issuer string
	var name string

	if claims["client_id"] != nil {
		companyID = claims["client_id"].(string)
	}

	if claims["sub"] != nil {
		sub = claims["sub"].(string)
	}

	if claims["issuer"] != nil {
		issuer = claims["issuer"].(string)
	}

	if claims["name"] != nil {
		name = claims["name"].(string)
	}

	if sub == "" {
		slog.Error("GetAuthContext - missing claims", "Error", ErrNoTokenFound)

		return AuthContext{}, ErrNoTokenFound
	}

	return AuthContext{
		Sub:       sub,
		Issuer:    issuer,
		Name:      name,
		CompanyID: companyID,
	}, nil
}
