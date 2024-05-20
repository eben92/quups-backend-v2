// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repository

import (
	"context"
	"database/sql"
)

type Querier interface {
	AddMember(ctx context.Context, arg AddMemberParams) (Member, error)
	// WORKING HOURS
	AddWorkingHour(ctx context.Context, arg AddWorkingHourParams) (WorkingHour, error)
	CreateCompany(ctx context.Context, arg CreateCompanyParams) (Company, error)
	CreateConfiguration(ctx context.Context, arg CreateConfigurationParams) (Configuration, error)
	CreatePaymentAccount(ctx context.Context, arg CreatePaymentAccountParams) (PaymentAccount, error)
	// PAYMENT ACCOUNT DETAILS
	CreatePaymentAccountDetails(ctx context.Context, arg CreatePaymentAccountDetailsParams) (PaymentAccountDetail, error)
	// PAYOUT ACCOUNT
	CreatePayoutAccount(ctx context.Context, arg CreatePayoutAccountParams) (PayoutAccount, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	// WARNING: this will not work because of foreign key constraints
	DeleteCompany(ctx context.Context, id string) error
	GetAllCompanies(ctx context.Context) ([]Company, error)
	GetCompanyByID(ctx context.Context, id string) (Company, error)
	GetCompanyByName(ctx context.Context, name string) (Company, error)
	GetConfigurationByCompanyID(ctx context.Context, companyID string) (Configuration, error)
	GetMembersByCompanyID(ctx context.Context, arg GetMembersByCompanyIDParams) ([]Member, error)
	GetPaymentAccountByCompanyID(ctx context.Context, companyID string) (PaymentAccount, error)
	GetPaymentAccountDetailsByPaymentAccountID(ctx context.Context, paymentAccountID string) (PaymentAccountDetail, error)
	GetPayoutAccountByPaymentAccountID(ctx context.Context, paymentAccountID string) (PayoutAccount, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id string) (User, error)
	GetUserByMsisdn(ctx context.Context, msisdn sql.NullString) (User, error)
	GetUserTeams(ctx context.Context, userID sql.NullString) ([]GetUserTeamsRow, error)
	GetUsers(ctx context.Context) ([]User, error)
	GetWorkingHoursByCompanyID(ctx context.Context, companyID string) ([]WorkingHour, error)
	UpdateCompany(ctx context.Context, arg UpdateCompanyParams) (Company, error)
	UpdateConfiguationByCompanyID(ctx context.Context, arg UpdateConfiguationByCompanyIDParams) (Configuration, error)
	UpdateMember(ctx context.Context, arg UpdateMemberParams) (Member, error)
	UpdateMemberInvitationStatus(ctx context.Context, arg UpdateMemberInvitationStatusParams) (Member, error)
	UpdatePaymentAccountByCompanyID(ctx context.Context, arg UpdatePaymentAccountByCompanyIDParams) (PaymentAccount, error)
	UpdatePaymentAccountDetailsByPaymentAccountID(ctx context.Context, arg UpdatePaymentAccountDetailsByPaymentAccountIDParams) (PaymentAccountDetail, error)
	UpdatePayoutAccountByPaymentAccountID(ctx context.Context, arg UpdatePayoutAccountByPaymentAccountIDParams) (PayoutAccount, error)
	UpdateWorkingHourByCompanyID(ctx context.Context, arg UpdateWorkingHourByCompanyIDParams) (WorkingHour, error)
}

var _ Querier = (*Queries)(nil)
