package paymentdto

import "time"

type Bank struct {
	Active           bool        `json:"active"`
	Code             string      `json:"code"`
	Country          string      `json:"country"`
	CreatedAt        time.Time   `json:"createdAt"`
	Currency         string      `json:"currency"`
	Gateway          interface{} `json:"gateway"`
	ID               int         `json:"id"`
	IsDeleted        bool        `json:"is_deleted"`
	Longcode         string      `json:"longcode"`
	Name             string      `json:"name"`
	PayWithBank      bool        `json:"pay_with_bank"`
	Slug             string      `json:"slug"`
	SupportsTransfer bool        `json:"supports_transfer"`
	Type             string      `json:"type"`
	UpdatedAt        time.Time   `json:"updatedAt"`
}

type ResolvedAccount struct {
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
	BankID        int    `json:"bank_id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
}
