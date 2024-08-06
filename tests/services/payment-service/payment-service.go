package paymentservice_test

import (
	"quups-backend/internal/database"
	paymentdto "quups-backend/internal/services/payment-service/dto"
	"quups-backend/internal/services/payment-service/models"
	paymentservice "quups-backend/internal/services/payment-service/service"
	local_jwt "quups-backend/internal/utils/jwt"
	testdb "quups-backend/tests/db"
	jwt_test "quups-backend/tests/utils/jwt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetupAccount_Integration(t *testing.T) {
	dbm := testdb.NewTestDatabase(t)
	defer dbm.Close(t)

	connstr := dbm.ConnectionString(t)

	db := database.NewService(connstr)

	ctx, err := jwt_test.SetTestContext(local_jwt.COMPANY_CTX_KEY)

	require.NoError(t, err, "error setting up test context")

	psvc := paymentservice.NewPaymentService(ctx, db)

	psvc.SetupAccount(paymentdto.ReqPaymentDTO{
		PaymentDetails: paymentdto.ReqBankDetailsDTO{
			BankCode:      "011",
			AccountNumber: "1234567890",
			BankID:        1,
			BankCurrency:  "GHS",
			BankName:      "GCB",
			FirstName:     "John",
			LastName:      "Doe",
			BankType:      string(models.BANK),
		},
		Address: paymentdto.ReqBillingAddressDTO{
			TIN:              "1234567890",
			Address:          "Accra",
			City:             "Accra",
			Country:          "Ghana",
			PostalCode:       "00233",
			Region:           "Greater Accra",
			Latitude:         5.6037,
			Longitude:        -0.1870,
			FormattedAddress: "Accra, Greater Accra, Ghana",
		},
		CompanyID: "1",
	})

}
