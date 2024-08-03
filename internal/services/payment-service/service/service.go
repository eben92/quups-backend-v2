package paymentservice

import (
	"context"
	"quups-backend/internal/database"
	paymentdto "quups-backend/internal/services/payment-service/dto"
	"quups-backend/internal/services/payment-service/models"
)

type service struct {
	db  database.Service
	ctx context.Context
}

// PaymentService provides methods for interacting with payment services.
type PaymentService interface {
	// GetBankList returns a list of supported banks.
	GetBankList(bankType models.BankType) ([]paymentdto.Bank, error)
	ResolveBankAccount(bankCode, accountNumber string) (paymentdto.ResolvedAccount, error)
}

// Payment service
func NewPaymentService(c context.Context, db database.Service) PaymentService {
	return &service{
		db:  db,
		ctx: c,
	}
}
