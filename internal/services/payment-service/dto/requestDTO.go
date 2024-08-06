package paymentdto

import "errors"

type ReqBankDetailsDTO struct {
	BankCode      string `json:"bank_code"`
	AccountNumber string `json:"account_number"`
	BankID        int32  `json:"bank_id"`
	BankCurrency  string `json:"bank_currency"`
	BankName      string `json:"bank_name"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	BankType      string `json:"bank_type"`
}

type ReqBillingAddressDTO struct {
	TIN              string  `json:"tin"`
	Address          string  `json:"address"`
	City             string  `json:"city"`
	Country          string  `json:"country"`
	PostalCode       string  `json:"postal_code"`
	Region           string  `json:"region"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	FormattedAddress string  `json:"formatted_address"`
}

type ReqPaymentDTO struct {
	PaymentDetails ReqBankDetailsDTO    `json:"payment_details"`
	Address        ReqBillingAddressDTO `json:"address"`
	CompanyID      string               `json:"company_id"`
}

func ValidateReqPaymentDTO(body ReqPaymentDTO) error {
	if body.PaymentDetails.BankCode == "" || body.PaymentDetails.AccountNumber == "" {
		return errors.New("bank_code and account_number are required")
	}

	if body.PaymentDetails.BankID == 0 || body.PaymentDetails.BankCurrency == "" || body.PaymentDetails.BankName == "" || body.PaymentDetails.FirstName == "" || body.PaymentDetails.LastName == "" {
		return errors.New("bank_id, bank_currency, bank_name, first_name and last_name are required")
	}

	if body.Address.TIN == "" || body.Address.Address == "" || body.Address.City == "" || body.Address.Country == "" || body.Address.PostalCode == "" || body.Address.Region == "" {
		return errors.New("tin, address, city, country, postal_code and region are required")
	}

	if body.CompanyID == "" {
		return errors.New("company_id is required")
	}

	return nil
}
