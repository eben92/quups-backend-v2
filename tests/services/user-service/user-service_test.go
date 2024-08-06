package userservice_test

import (
	"context"
	"quups-backend/internal/database"
	userdto "quups-backend/internal/services/user-service/dto"
	userservice "quups-backend/internal/services/user-service/service"
	testdb "quups-backend/tests/db"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser_Integration(t *testing.T) {
	db := testdb.NewTestDatabase(t)
	defer db.Close(t)
	connStr := db.ConnectionString(t)
	dbM := database.NewService(connStr)

	usrv := userservice.NewUserService(context.Background(), dbM)

	user, err := usrv.Create(userdto.CreateUserParams{
		Email:    "test@email.com",
		Name:     "Test User",
		Msisdn:   "0241234567",
		Gender:   "Male",
		Password: "1123445",
	})

	require.NoError(t, err, "Error creating user")

	require.NotEmpty(t, user.ID)
	require.NotEmpty(t, user.Email)
	require.NotEmpty(t, user.Name)
	require.NotEmpty(t, user.Msisdn)
	require.NotEmpty(t, user.Gender)
	require.NotEmpty(t, user.Password)

}
