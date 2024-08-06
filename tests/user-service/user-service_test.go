package userservice_test

import (
	"context"
	"database/sql"
	"quups-backend/internal/database"
	"quups-backend/internal/database/repository"
	testdb "quups-backend/tests/db"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateUser_Integration(t *testing.T) {
	db := testdb.NewTestDatabase(t)
	defer db.Close(t)
	connStr := db.ConnectionString(t)
	dbM := database.NewService(connStr)

	// usrv := userservice.NewUserService(context.Background(), dbM)
	repo := dbM.NewRepository()

	_, err := repo.CreateUser(context.Background(), repository.CreateUserParams{
		Email: "test@tes.com",
		Name: sql.NullString{
			String: "Test User", Valid: true},
		Msisdn:   sql.NullString{String: "0241234567", Valid: true},
		ImageUrl: sql.NullString{String: "dfdfdf", Valid: true},
		Gender:   sql.NullString{String: "male", Valid: true},
		Dob:      sql.NullTime{Time: time.Now(), Valid: true},
		Otp:      sql.NullString{String: "1234", Valid: true},
		Password: sql.NullString{String: "1234", Valid: true},
	})

	// user, err := usrv.Create(userdto.CreateUserParams{
	// 	Email:    "test@email.com",
	// 	Name:     "Test User",
	// 	Msisdn:   "0241234567",
	// 	Gender:   "Male",
	// 	Password: "1123445",
	// })

	require.NoError(t, err, "Error creating user")

	// require.NotEmpty(t, user.ID)
	// require.NotEmpty(t, user.Email)
	// require.NotEmpty(t, user.Name)
	// require.NotEmpty(t, user.Msisdn)
	// require.NotEmpty(t, user.Gender)
	// require.NotEmpty(t, user.Password)

}
