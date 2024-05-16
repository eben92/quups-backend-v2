package local_jwt

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	COOKIE_NAME string = "qp_session"
)

var (
	JWT_SECRET   = os.Getenv("JWT_SECRET")
	AUTH_CTX_KEY = &contextKey{"authcontext"}
)

type authContext struct {
	Sub string
}

type contextKey struct {
	name string
}

var (
	ErrUnauthorized = errors.New("token is unauthorized")
	ErrExpired      = errors.New("token is expired")
	ErrNBFInvalid   = errors.New("token nbf validation failed")
	ErrIATInvalid   = errors.New("token iat validation failed")
	ErrNoTokenFound = errors.New("no token found")
	ErrAlgoInvalid  = errors.New("algorithm mismatch")
)

// Authenticator is a middleware that handles jwt authentications
func Authenticator() func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		hfn := func(w http.ResponseWriter, r *http.Request) {

			var token string

			findtokens := []func(*http.Request) string{GetTokenFromHeader, GetTokenFromCookie}

			for _, fn := range findtokens {
				token = fn(r)

				if token != "" {
					break
				}

			}

			if token == "" {

				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return

			}

			c, err := ParseToken(token)

			if err != nil {

				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			ctx, err := newContext(r.Context(), c)

			if err != nil {

				http.Error(w, err.Error(), http.StatusUnauthorized)
				return

			}

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(hfn)
	}

}

func GetTokenFromHeader(r *http.Request) string {
	bearer := r.Header.Get("Authorization")

	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		log.Println("got token from request head")
		return bearer[7:]
	}

	return ""
}

func GetTokenFromCookie(r *http.Request) string {

	cookie, err := r.Cookie(COOKIE_NAME)

	if err != nil {
		return ""
	}

	log.Printf("token found in cookie")

	return cookie.Value

}

/*
Generetes a signed token and return as byte or nil.
Convert to string before sending to client
*/
func GenereteJWT(ID, name string) ([]byte, error) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    ID,
		"issuer": "WEB",
		"name":   name,
		"exp":    time.Now().Add(time.Minute).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(JWT_SECRET))

	if err != nil {
		log.Printf("Error signing jwt [%s]", err.Error())

		return nil, fmt.Errorf("Something went wrong. Please try again. #2")
	}

	return []byte(tokenString), nil
}

// ParseToken is use to verify jwt tokens
func ParseToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, alg := token.Method.(*jwt.SigningMethodHMAC); !alg {

			log.Printf("invalid token alg [%s]", token.Header["alg"])

			return nil, ErrAlgoInvalid
		}

		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		log.Printf("error parsing token [%s]", err.Error())

		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, ErrNBFInvalid
	}

	return claims, nil

}

func newContext(ctx context.Context, claims jwt.MapClaims) (context.Context, error) {
	sub := claims["sub"]

	if sub == "" {
		return nil, ErrUnauthorized
	}

	ctx = context.WithValue(ctx, AUTH_CTX_KEY, claims)
	return ctx, nil
}

// GetAuthContext returns decoded jwt data
func GetAuthContext(ctx context.Context) *authContext {
	claims, _ := ctx.Value(AUTH_CTX_KEY).(jwt.MapClaims)

	return &authContext{
		Sub: claims["sub"].(string),
	}

}
