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
	db := testdb.NewDatabase(t)
	defer db.Close(t)
	connStr := db.ConnectionString(t)
	dbM := database.NewService(connStr)

	columns := []string{
		"id SERIAL PRIMARY KEY",
		"email VARCHAR(100) NOT NULL",
		"name VARCHAR(100) NOT NULL",
		"msisdn VARCHAR(100) NOT NULL",
		"gender VARCHAR(100)",
		"image_url VARCHAR(100) ",
		"password VARCHAR(100) NOT NULL",
		"tin_number VARCHAR(100)",
		"otp VARCHAR(100)",
		"email_verified TIMESTAMP WITH TIME ZONE",
		"app_push_token VARCHAR(150)",
		"web_push_token VARCHAR(150)",
		"dob TIMESTAMP WITHOUT TIME ZONE",
		"created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP",
		"updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP",
	}
	err := dbM.CreateTable("users", columns)
	require.NoError(t, err, "Error creating table")

	usrv := userservice.NewUserService(context.TODO(), dbM)

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

}
