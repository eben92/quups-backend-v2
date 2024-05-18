// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repository

import (
	"database/sql"
	"time"
)

type Account struct {
	ID                string         `json:"id"`
	Provider          string         `json:"provider"`
	ProviderAccountID string         `json:"provider_account_id"`
	Type              string         `json:"type"`
	ExpiresAt         int32          `json:"expires_at"`
	TokenType         string         `json:"token_type"`
	AccessToken       sql.NullString `json:"access_token"`
	RefreshToken      sql.NullString `json:"refresh_token"`
	AccountType       sql.NullString `json:"account_type"`
	IDToken           sql.NullString `json:"id_token"`
	Scope             sql.NullString `json:"scope"`
	UserID            string         `json:"user_id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
}

type Address struct {
	ID               string          `json:"id"`
	Street           string          `json:"street"`
	City             string          `json:"city"`
	Region           sql.NullString  `json:"region"`
	Country          string          `json:"country"`
	CountryCode      string          `json:"country_code"`
	FormattedAddress sql.NullString  `json:"formatted_address"`
	Description      sql.NullString  `json:"description"`
	PostalCode       sql.NullString  `json:"postal_code"`
	Latitude         sql.NullFloat64 `json:"latitude"`
	Longitude        sql.NullFloat64 `json:"longitude"`
	Msisdn           sql.NullString  `json:"msisdn"`
	IsDefault        sql.NullBool    `json:"is_default"`
	UserID           sql.NullString  `json:"user_id"`
	CompanyID        sql.NullString  `json:"company_id"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}

type Company struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Slug           string         `json:"slug"`
	About          sql.NullString `json:"about"`
	Msisdn         string         `json:"msisdn"`
	Email          string         `json:"email"`
	Tin            sql.NullString `json:"tin"`
	ImageUrl       sql.NullString `json:"image_url"`
	BannerUrl      sql.NullString `json:"banner_url"`
	BrandType      string         `json:"brand_type"`
	OwnerID        string         `json:"owner_id"`
	TotalSales     int32          `json:"total_sales"`
	IsActive       bool           `json:"is_active"`
	CurrencyCode   string         `json:"currency_code"`
	InvitationCode sql.NullString `json:"invitation_code"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

type CompanyEmployee struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Msisdn    string         `json:"msisdn"`
	Email     sql.NullString `json:"email"`
	Role      string         `json:"role"`
	Status    string         `json:"status"`
	CompanyID string         `json:"company_id"`
	UserID    sql.NullString `json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type Configuration struct {
	ID              string    `json:"id"`
	Delivery        bool      `json:"delivery"`
	Pickup          bool      `json:"pickup"`
	CashOnDelivery  bool      `json:"cash_on_delivery"`
	DigitalPayments bool      `json:"digital_payments"`
	CompanyID       string    `json:"company_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type PaymentAccount struct {
	ID            string    `json:"id"`
	AccountType   string    `json:"account_type"`
	AccountNumber string    `json:"account_number"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	BankName      string    `json:"bank_name"`
	BankCode      string    `json:"bank_code"`
	BankBranch    string    `json:"bank_branch"`
	CompanyID     string    `json:"company_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type PaymentAccountDetail struct {
	ID               string         `json:"id"`
	IDInt            int32          `json:"id_int"`
	Name             string         `json:"name"`
	Slug             string         `json:"slug"`
	Code             string         `json:"code"`
	Longcode         sql.NullString `json:"longcode"`
	Gateway          sql.NullString `json:"gateway"`
	PayWithBank      bool           `json:"pay_with_bank"`
	Active           bool           `json:"active"`
	IsDeleted        bool           `json:"is_deleted"`
	Country          string         `json:"country"`
	Currency         string         `json:"currency"`
	Type             string         `json:"type"`
	PaymentAccountID string         `json:"payment_account_id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}

type PayoutAccount struct {
	ID                 string         `json:"id"`
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
	PaymentAccountID   string         `json:"payment_account_id"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
}

type User struct {
	ID            string         `json:"id"`
	Email         string         `json:"email"`
	Msisdn        sql.NullString `json:"msisdn"`
	EmailVerified sql.NullTime   `json:"email_verified"`
	Name          sql.NullString `json:"name"`
	ImageUrl      sql.NullString `json:"image_url"`
	TinNumber     sql.NullString `json:"tin_number"`
	Gender        sql.NullString `json:"gender"`
	Dob           sql.NullTime   `json:"dob"`
	Otp           sql.NullString `json:"otp"`
	AppPushToken  sql.NullString `json:"app_push_token"`
	WebPushToken  sql.NullString `json:"web_push_token"`
	Password      sql.NullString `json:"password"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

type WorkingHour struct {
	ID        string    `json:"id"`
	Day       string    `json:"day"`
	OpensAt   time.Time `json:"opens_at"`
	ClosesAt  time.Time `json:"closes_at"`
	CompanyID string    `json:"company_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
