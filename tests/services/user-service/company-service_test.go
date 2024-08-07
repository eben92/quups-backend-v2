package userservice_test

import (
	"context"
	"quups-backend/internal/database"
	paymentdto "quups-backend/internal/services/payment-service/dto"
	local_jwt "quups-backend/internal/utils/jwt"
	testdb "quups-backend/tests/db"
	mock_test "quups-backend/tests/mock"
	jwt_test "quups-backend/tests/utils/jwt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetupAccount_Integration(t *testing.T) {
	dbm := testdb.NewTestDatabase(t)
	defer dbm.Close(t)

	connStr := dbm.ConnectionString(t)

	db := database.NewService(connStr)

	mkusvc := mock_test.NewMockSvc(context.Background(), db)

	sampleUsr := mock_test.GetSampleUser()
	usr, err := mkusvc.CreateUser(sampleUsr)

	require.NoError(t, err, "error creating user")
	require.NotEmpty(t, usr.ID, "user ID is empty")

	usrCtx, err := jwt_test.SetTestContext(local_jwt.AUTH_CTX_KEY, jwt_test.Claims{
		Sub:    usr.ID,
		Issuer: "test",
		Name:   usr.Name,
		// ClientID: ,
	})

	require.NoError(t, err, "error setting test context")

	mkCompSvc := mock_test.NewMockCompanySvc(usrCtx, db)

	comp, err := mkCompSvc.CreateCompany(mock_test.GetSampleCompany())

	require.NoError(t, err, "error creating company")
	require.NotEmpty(t, comp.ID, "company ID is empty")

	compCtx, err := jwt_test.SetTestContext(local_jwt.COMPANY_CTX_KEY, jwt_test.Claims{
		Sub:      comp.ID,
		Issuer:   "test",
		Name:     comp.Name,
		ClientID: comp.ID,
	})

	require.NoError(t, err, "error setting test context")

	paymSvc := mock_test.NewMockPaymentSvc(compCtx, db)

	sampAddr := mock_test.GetSampleBillingAddress()
	sampBankDts := mock_test.GetSampleBankDetails()

	err = paymSvc.SetupAccount(paymentdto.ReqPaymentDTO{
		PaymentDetails: sampBankDts,
		Address:        sampAddr,
	})

	require.NoError(t, err, "error setting up account")

}
