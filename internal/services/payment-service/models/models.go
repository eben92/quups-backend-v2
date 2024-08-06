package models

import "time"

type BankType string

const (
	MOBILE_MONEY BankType = "mobile_money"
	BANK         BankType = "ghipss"
	CHARGE       int      = 8
)

type ThirdPartyWallet struct {
	BusinessName        string    `json:"business_name"`
	PrimaryContactName  string    `json:"primary_contact_name"`
	PrimaryContactEmail string    `json:"primary_contact_email"`
	PrimaryContactPhone string    `json:"primary_contact_phone"`
	AccountNumber       string    `json:"account_number"`
	PercentageCharge    float64   `json:"percentage_charge"`
	SettlementBank      string    `json:"settlement_bank"`
	Currency            string    `json:"currency"`
	Bank                int       `json:"bank"`
	SubAccountCode      string    `json:"subaccount_code"`
	Integration         int       `json:"integration"`
	Domain              string    `json:"domain"`
	IsVerified          bool      `json:"is_verified"`
	Active              bool      `json:"active"`
	ID                  int       `json:"id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
