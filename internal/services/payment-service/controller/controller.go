package paymentcontroller

import (
	"net/http"
	"quups-backend/internal/database"
)

type controller struct {
	db database.Service
}

type paymentController interface {
	GetBankList(w http.ResponseWriter, r *http.Request)
	ResolveBankAccount(w http.ResponseWriter, r *http.Request)
}

func NewPaymentController(db database.Service) paymentController {
	return &controller{
		db: db,
	}
}
