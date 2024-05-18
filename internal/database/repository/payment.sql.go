// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: payment.sql

package repository

import (
	"context"
	"database/sql"
)

const createPaymentAccount = `-- name: CreatePaymentAccount :one
INSERT INTO payment_accounts (
        company_id,
        account_number,
        account_type,
        first_name,
        last_name,
        bank_branch,
        bank_code,
        bank_name
    )
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
    )
RETURNING id, account_type, account_number, first_name, last_name, bank_name, bank_code, bank_branch, company_id, created_at, updated_at
`

type CreatePaymentAccountParams struct {
	CompanyID     string `json:"company_id"`
	AccountNumber string `json:"account_number"`
	AccountType   string `json:"account_type"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	BankBranch    string `json:"bank_branch"`
	BankCode      string `json:"bank_code"`
	BankName      string `json:"bank_name"`
}

func (q *Queries) CreatePaymentAccount(ctx context.Context, arg CreatePaymentAccountParams) (PaymentAccount, error) {
	row := q.db.QueryRowContext(ctx, createPaymentAccount,
		arg.CompanyID,
		arg.AccountNumber,
		arg.AccountType,
		arg.FirstName,
		arg.LastName,
		arg.BankBranch,
		arg.BankCode,
		arg.BankName,
	)
	var i PaymentAccount
	err := row.Scan(
		&i.ID,
		&i.AccountType,
		&i.AccountNumber,
		&i.FirstName,
		&i.LastName,
		&i.BankName,
		&i.BankCode,
		&i.BankBranch,
		&i.CompanyID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createPaymentAccountDetails = `-- name: CreatePaymentAccountDetails :one

INSERT INTO payment_account_details (
        payment_account_id,
        id_int,
        currency,
        name,
        slug,
        code,
        longcode,
        gateway,
        pay_with_bank,
        is_deleted,
        country,
        type,
        active
    )
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
    )
RETURNING id, id_int, name, slug, code, longcode, gateway, pay_with_bank, active, is_deleted, country, currency, type, payment_account_id, created_at, updated_at
`

type CreatePaymentAccountDetailsParams struct {
	PaymentAccountID string         `json:"payment_account_id"`
	IDInt            int32          `json:"id_int"`
	Currency         string         `json:"currency"`
	Name             string         `json:"name"`
	Slug             string         `json:"slug"`
	Code             string         `json:"code"`
	Longcode         sql.NullString `json:"longcode"`
	Gateway          sql.NullString `json:"gateway"`
	PayWithBank      bool           `json:"pay_with_bank"`
	IsDeleted        bool           `json:"is_deleted"`
	Country          string         `json:"country"`
	Type             string         `json:"type"`
	Active           bool           `json:"active"`
}

// PAYMENT ACCOUNT DETAILS
func (q *Queries) CreatePaymentAccountDetails(ctx context.Context, arg CreatePaymentAccountDetailsParams) (PaymentAccountDetail, error) {
	row := q.db.QueryRowContext(ctx, createPaymentAccountDetails,
		arg.PaymentAccountID,
		arg.IDInt,
		arg.Currency,
		arg.Name,
		arg.Slug,
		arg.Code,
		arg.Longcode,
		arg.Gateway,
		arg.PayWithBank,
		arg.IsDeleted,
		arg.Country,
		arg.Type,
		arg.Active,
	)
	var i PaymentAccountDetail
	err := row.Scan(
		&i.ID,
		&i.IDInt,
		&i.Name,
		&i.Slug,
		&i.Code,
		&i.Longcode,
		&i.Gateway,
		&i.PayWithBank,
		&i.Active,
		&i.IsDeleted,
		&i.Country,
		&i.Currency,
		&i.Type,
		&i.PaymentAccountID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createPayoutAccount = `-- name: CreatePayoutAccount :one

INSERT INTO payout_accounts (
        payment_account_id,
        id_int,
        currency,
        business_name,
        account_number,
        primay_contact_name,
        primay_contact_email,
        primay_contact_phone,
        description,
        subaccount_code,
        settlement_bank,
        percentage_charge,
        active,
        bank_id
    )
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
    )
RETURNING id, id_int, currency, business_name, account_number, primay_contact_name, primay_contact_email, primay_contact_phone, description, subaccount_code, settlement_bank, percentage_charge, active, bank_id, payment_account_id, created_at, updated_at
`

type CreatePayoutAccountParams struct {
	PaymentAccountID   string         `json:"payment_account_id"`
	IDInt              int32          `json:"id_int"`
	Currency           string         `json:"currency"`
	BusinessName       string         `json:"business_name"`
	AccountNumber      string         `json:"account_number"`
	PrimayContactName  string         `json:"primay_contact_name"`
	PrimayContactEmail string         `json:"primay_contact_email"`
	PrimayContactPhone string         `json:"primay_contact_phone"`
	Description        sql.NullString `json:"description"`
	SubaccountCode     string         `json:"subaccount_code"`
	SettlementBank     string         `json:"settlement_bank"`
	PercentageCharge   float64        `json:"percentage_charge"`
	Active             bool           `json:"active"`
	BankID             string         `json:"bank_id"`
}

// PAYOUT ACCOUNT
func (q *Queries) CreatePayoutAccount(ctx context.Context, arg CreatePayoutAccountParams) (PayoutAccount, error) {
	row := q.db.QueryRowContext(ctx, createPayoutAccount,
		arg.PaymentAccountID,
		arg.IDInt,
		arg.Currency,
		arg.BusinessName,
		arg.AccountNumber,
		arg.PrimayContactName,
		arg.PrimayContactEmail,
		arg.PrimayContactPhone,
		arg.Description,
		arg.SubaccountCode,
		arg.SettlementBank,
		arg.PercentageCharge,
		arg.Active,
		arg.BankID,
	)
	var i PayoutAccount
	err := row.Scan(
		&i.ID,
		&i.IDInt,
		&i.Currency,
		&i.BusinessName,
		&i.AccountNumber,
		&i.PrimayContactName,
		&i.PrimayContactEmail,
		&i.PrimayContactPhone,
		&i.Description,
		&i.SubaccountCode,
		&i.SettlementBank,
		&i.PercentageCharge,
		&i.Active,
		&i.BankID,
		&i.PaymentAccountID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPaymentAccountByCompanyID = `-- name: GetPaymentAccountByCompanyID :one
SELECT id, account_type, account_number, first_name, last_name, bank_name, bank_code, bank_branch, company_id, created_at, updated_at FROM payment_accounts
 WHERE company_id = $1
 LIMIT 1
`

func (q *Queries) GetPaymentAccountByCompanyID(ctx context.Context, companyID string) (PaymentAccount, error) {
	row := q.db.QueryRowContext(ctx, getPaymentAccountByCompanyID, companyID)
	var i PaymentAccount
	err := row.Scan(
		&i.ID,
		&i.AccountType,
		&i.AccountNumber,
		&i.FirstName,
		&i.LastName,
		&i.BankName,
		&i.BankCode,
		&i.BankBranch,
		&i.CompanyID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPaymentAccountDetailsByPaymentAccountID = `-- name: GetPaymentAccountDetailsByPaymentAccountID :one
SELECT id, id_int, name, slug, code, longcode, gateway, pay_with_bank, active, is_deleted, country, currency, type, payment_account_id, created_at, updated_at FROM payment_account_details
 WHERE payment_account_id = $1
 LIMIT 1
`

func (q *Queries) GetPaymentAccountDetailsByPaymentAccountID(ctx context.Context, paymentAccountID string) (PaymentAccountDetail, error) {
	row := q.db.QueryRowContext(ctx, getPaymentAccountDetailsByPaymentAccountID, paymentAccountID)
	var i PaymentAccountDetail
	err := row.Scan(
		&i.ID,
		&i.IDInt,
		&i.Name,
		&i.Slug,
		&i.Code,
		&i.Longcode,
		&i.Gateway,
		&i.PayWithBank,
		&i.Active,
		&i.IsDeleted,
		&i.Country,
		&i.Currency,
		&i.Type,
		&i.PaymentAccountID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPayoutAccountByPaymentAccountID = `-- name: GetPayoutAccountByPaymentAccountID :one
SELECT id, id_int, currency, business_name, account_number, primay_contact_name, primay_contact_email, primay_contact_phone, description, subaccount_code, settlement_bank, percentage_charge, active, bank_id, payment_account_id, created_at, updated_at FROM payout_accounts
 WHERE payment_account_id = $1
 LIMIT 1
`

func (q *Queries) GetPayoutAccountByPaymentAccountID(ctx context.Context, paymentAccountID string) (PayoutAccount, error) {
	row := q.db.QueryRowContext(ctx, getPayoutAccountByPaymentAccountID, paymentAccountID)
	var i PayoutAccount
	err := row.Scan(
		&i.ID,
		&i.IDInt,
		&i.Currency,
		&i.BusinessName,
		&i.AccountNumber,
		&i.PrimayContactName,
		&i.PrimayContactEmail,
		&i.PrimayContactPhone,
		&i.Description,
		&i.SubaccountCode,
		&i.SettlementBank,
		&i.PercentageCharge,
		&i.Active,
		&i.BankID,
		&i.PaymentAccountID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePaymentAccountByCompanyID = `-- name: UpdatePaymentAccountByCompanyID :one
UPDATE payment_accounts SET 
        account_number = $2,
        account_type = $3,
        first_name = $4,
        last_name = $5,
        bank_branch = $6,
        bank_code = $7,
        bank_name = $8
    WHERE company_id = $1
RETURNING id, account_type, account_number, first_name, last_name, bank_name, bank_code, bank_branch, company_id, created_at, updated_at
`

type UpdatePaymentAccountByCompanyIDParams struct {
	CompanyID     string `json:"company_id"`
	AccountNumber string `json:"account_number"`
	AccountType   string `json:"account_type"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	BankBranch    string `json:"bank_branch"`
	BankCode      string `json:"bank_code"`
	BankName      string `json:"bank_name"`
}

func (q *Queries) UpdatePaymentAccountByCompanyID(ctx context.Context, arg UpdatePaymentAccountByCompanyIDParams) (PaymentAccount, error) {
	row := q.db.QueryRowContext(ctx, updatePaymentAccountByCompanyID,
		arg.CompanyID,
		arg.AccountNumber,
		arg.AccountType,
		arg.FirstName,
		arg.LastName,
		arg.BankBranch,
		arg.BankCode,
		arg.BankName,
	)
	var i PaymentAccount
	err := row.Scan(
		&i.ID,
		&i.AccountType,
		&i.AccountNumber,
		&i.FirstName,
		&i.LastName,
		&i.BankName,
		&i.BankCode,
		&i.BankBranch,
		&i.CompanyID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePaymentAccountDetailsByPaymentAccountID = `-- name: UpdatePaymentAccountDetailsByPaymentAccountID :one
UPDATE payment_account_details SET 
        name = $2,
        active = $3,
        currency = $4,
        type = $5,
        slug = $6,
        code = $7
        
    WHERE payment_account_id = $1
RETURNING id, id_int, name, slug, code, longcode, gateway, pay_with_bank, active, is_deleted, country, currency, type, payment_account_id, created_at, updated_at
`

type UpdatePaymentAccountDetailsByPaymentAccountIDParams struct {
	PaymentAccountID string `json:"payment_account_id"`
	Name             string `json:"name"`
	Active           bool   `json:"active"`
	Currency         string `json:"currency"`
	Type             string `json:"type"`
	Slug             string `json:"slug"`
	Code             string `json:"code"`
}

func (q *Queries) UpdatePaymentAccountDetailsByPaymentAccountID(ctx context.Context, arg UpdatePaymentAccountDetailsByPaymentAccountIDParams) (PaymentAccountDetail, error) {
	row := q.db.QueryRowContext(ctx, updatePaymentAccountDetailsByPaymentAccountID,
		arg.PaymentAccountID,
		arg.Name,
		arg.Active,
		arg.Currency,
		arg.Type,
		arg.Slug,
		arg.Code,
	)
	var i PaymentAccountDetail
	err := row.Scan(
		&i.ID,
		&i.IDInt,
		&i.Name,
		&i.Slug,
		&i.Code,
		&i.Longcode,
		&i.Gateway,
		&i.PayWithBank,
		&i.Active,
		&i.IsDeleted,
		&i.Country,
		&i.Currency,
		&i.Type,
		&i.PaymentAccountID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePayoutAccountByPaymentAccountID = `-- name: UpdatePayoutAccountByPaymentAccountID :one
UPDATE payout_accounts SET 
        percentage_charge = $2,
        active = $3,
        account_number = $4,
        subaccount_code = $5,
        settlement_bank = $6

    WHERE payment_account_id = $1
RETURNING id, id_int, currency, business_name, account_number, primay_contact_name, primay_contact_email, primay_contact_phone, description, subaccount_code, settlement_bank, percentage_charge, active, bank_id, payment_account_id, created_at, updated_at
`

type UpdatePayoutAccountByPaymentAccountIDParams struct {
	PaymentAccountID string  `json:"payment_account_id"`
	PercentageCharge float64 `json:"percentage_charge"`
	Active           bool    `json:"active"`
	AccountNumber    string  `json:"account_number"`
	SubaccountCode   string  `json:"subaccount_code"`
	SettlementBank   string  `json:"settlement_bank"`
}

func (q *Queries) UpdatePayoutAccountByPaymentAccountID(ctx context.Context, arg UpdatePayoutAccountByPaymentAccountIDParams) (PayoutAccount, error) {
	row := q.db.QueryRowContext(ctx, updatePayoutAccountByPaymentAccountID,
		arg.PaymentAccountID,
		arg.PercentageCharge,
		arg.Active,
		arg.AccountNumber,
		arg.SubaccountCode,
		arg.SettlementBank,
	)
	var i PayoutAccount
	err := row.Scan(
		&i.ID,
		&i.IDInt,
		&i.Currency,
		&i.BusinessName,
		&i.AccountNumber,
		&i.PrimayContactName,
		&i.PrimayContactEmail,
		&i.PrimayContactPhone,
		&i.Description,
		&i.SubaccountCode,
		&i.SettlementBank,
		&i.PercentageCharge,
		&i.Active,
		&i.BankID,
		&i.PaymentAccountID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
