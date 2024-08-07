package jwt_test

import (
	"context"
	local_jwt "quups-backend/internal/utils/jwt"

	"github.com/golang-jwt/jwt/v5"
)

func GetTestAuthContext() local_jwt.AuthContext {
	return local_jwt.AuthContext{
		CompanyID: "123456",
		Sub:       "test",
		Issuer:    "issuer",
		Name:      "test",
	}
}

type Claims struct {
	Sub      string `json:"sub"`
	Issuer   string `json:"issuer"`
	Name     string `json:"name"`
	ClientID string `json:"client_id"`
}

func SetTestContext(key *local_jwt.ContextKey, c Claims) (context.Context, error) {
	claims := jwt.MapClaims{
		"sub":       c.Sub,
		"issuer":    c.Issuer,
		"name":      c.Name,
		"client_id": c.ClientID,
	}

	ctx := context.Background()
	ctx, err := newContext(ctx, key, claims)
	if err != nil {
		return nil, err
	}

	return ctx, nil
}

func newContext(ctx context.Context, key *local_jwt.ContextKey, claims jwt.MapClaims) (context.Context, error) {
	sub := claims["sub"]

	if sub == "" {
		return nil, local_jwt.ErrUnauthorized
	}

	ctx = context.WithValue(ctx, key, claims)
	return ctx, nil
}
