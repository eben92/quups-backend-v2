package mock_test

import (
	"context"
	"quups-backend/internal/database"
	paymentdto "quups-backend/internal/services/payment-service/dto"
	"quups-backend/internal/services/payment-service/models"
	paymentservice "quups-backend/internal/services/payment-service/service"
)

type MockPaymentMgt interface {
}

type MockPaymentSvc struct {
	db  database.Service
	ctx context.Context
}

func NewMockPaymentSvc(ctx context.Context, db database.Service) MockPaymentMgt {
	return &MockPaymentSvc{db: db, ctx: ctx}
}

func GetSampleBillingAddress() paymentdto.ReqBillingAddressDTO {

	return paymentdto.ReqBillingAddressDTO{
		TIN:              "123456",
		Address:          "Test Address",
		City:             "Accra",
		Country:          "Ghana",
		PostalCode:       "00233",
		Region:           "Greater Accra",
		Latitude:         5.6037,
		Longitude:        -0.1870,
		FormattedAddress: "Test Address, Accra, Ghana",
	}

}

func GetSampleBankDetails() paymentdto.ReqBankDetailsDTO {
	return paymentdto.ReqBankDetailsDTO{
		AccountNumber: "1234567890",
		BankCode:      "123",
		BankID:        2,
		BankCurrency:  "GHS",
		BankName:      "Test Bank",
		FirstName:     "Test",
		LastName:      "User",
		BankType:      string(models.MOBILE_MONEY),
	}
}

func (m *MockPaymentSvc) SetupAccount(data paymentdto.ReqPaymentDTO) error {

	pmsvc := paymentservice.NewPaymentService(m.ctx, m.db)

	err := pmsvc.SetupAccount(data)

	return err
}
