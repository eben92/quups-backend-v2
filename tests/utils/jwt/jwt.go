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

func SetTestContext(key *local_jwt.ContextKey) (context.Context, error) {
	claims := jwt.MapClaims{
		"sub":       "test",
		"issuer":    "issuer",
		"name":      "test",
		"client_id": "123456",
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
