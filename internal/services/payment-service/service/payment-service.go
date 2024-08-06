package paymentservice

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"quups-backend/internal/database/repository"
	paymentdto "quups-backend/internal/services/payment-service/dto"
	"quups-backend/internal/services/payment-service/models"
	userdto "quups-backend/internal/services/user-service/dto"
	userservice "quups-backend/internal/services/user-service/service"
	local_http "quups-backend/internal/utils/http"
	"strings"
	"time"
)

var (
	PAYSTACK_URL string = os.Getenv("PAYSTACK_BASE_URL")
	ACC_TOKEN    string = os.Getenv("PAYSTACK_SECRET")
)

type bankListResponse struct {
	Status  bool
	Message string
	Data    []paymentdto.Bank
}

func (s *service) GetBankList(bankType models.BankType) ([]paymentdto.Bank, error) {
	var response bankListResponse

	url := fmt.Sprintf("%s/bank?country=ghana&type=%s", PAYSTACK_URL, bankType)
	bearer := fmt.Sprintf("Bearer %s", ACC_TOKEN)

	res, err := local_http.Fetch(url, &local_http.Options{
		Headers: &[][2]string{
			{"Authorization", bearer},
		},
	})

	if err != nil {

		slog.Error("GetBankList", "error", err)

		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		slog.Error("GetBankList - Decode", "error", err)

		return nil, err
	}

	return response.Data, nil

}

type resolveAccountResponse struct {
	Status  bool
	Message string
	Data    paymentdto.ResolvedAccount
}

func (s *service) ResolveBankAccount(bankCode, accountNumber string) (paymentdto.ResolvedAccount, error) {
	var response resolveAccountResponse
	var result paymentdto.ResolvedAccount

	// TODO: update this to 5
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	slog.Info("about to resolve payment account", "bankCode", bankCode, "accountNumber", accountNumber)

	url := fmt.Sprintf("%s/bank/resolve?account_number=%s&bank_code=%s", PAYSTACK_URL, accountNumber, bankCode)
	bearer := fmt.Sprintf("Bearer %s", ACC_TOKEN)

	res, err := local_http.FetchWithContext(ctx, url, &local_http.Options{
		Headers: &[][2]string{
			{"Authorization", bearer},
		},
	})

	if err != nil {
		slog.Error("ResolveBankAccount", "error", err)

		return result, fmt.Errorf("request  took longer than expected. please ensure you have a stable internet connection")
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		slog.Error("ResolveBankAccount - Decode", "error", err)

		return result, err
	}

	if !response.Status {
		slog.Error("ResolveBankAccount - response", "message", response.Message)

		return result, fmt.Errorf("failed to resolve account")
	}

	result = response.Data

	names := strings.Split(result.AccountName, " ")

	result.FirstName = names[0]
	if len(names) > 1 {
		result.LastName = strings.Join(names[1:], " ")
	}

	slog.Info("resolved account successfully", "result", result)

	return result, nil
}

type walletResponse struct {
	Status  bool
	Message string
	Data    models.ThirdPartyWallet
}

func (s *service) setupThirdPartyWallet(company userdto.CompanyInternalDTO, data paymentdto.ReqPaymentDTO) (models.ThirdPartyWallet, error) {
	var response walletResponse

	url := fmt.Sprintf("%s/subaccount", PAYSTACK_URL)
	bearer := fmt.Sprintf("Bearer %s", ACC_TOKEN)

	payload := map[string]interface{}{
		"business_name":         company.Name,
		"primary_contact_email": company.Email,
		"primary_contact_phone": company.Msisdn,
		"settlement_bank":       data.PaymentDetails.BankCode,
		"account_number":        data.PaymentDetails.AccountNumber,
		"percentage_charge":     models.CHARGE,
		"primary_contact_name":  data.PaymentDetails.FirstName + " " + data.PaymentDetails.LastName,
	}

	res, err := local_http.Fetch(url, &local_http.Options{
		Method: http.MethodPost,
		Body:   strings.NewReader(fmt.Sprintf("%v", payload)),
		Headers: &[][2]string{
			{"Authorization", bearer},
		},
	})

	if err != nil {
		slog.Error("setupThirdPartyWallet", "error", err)

		return response.Data, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		slog.Error("setupThirdPartyWallet - Decode", "error", err)

		return response.Data, err
	}

	return response.Data, nil
}

func (s *service) setUpPayoutAccont(acc models.ThirdPartyWallet) error {
	repo := s.db.NewRepository()

	_, err := repo.CreatePayoutAccount(s.ctx, repository.CreatePayoutAccountParams{})

	if err != nil {
		slog.Error("setUpPayoutAccont - CreatePayoutAccount", "error", err)

		return errors.New("failed to setup payout account. please try again")
	}

	return nil
}

func (s *service) AddBillingAddress(company userdto.CompanyInternalDTO, data paymentdto.ReqBillingAddressDTO) error {
	repo := s.db.NewRepository()
	formattedAddr := fmt.Sprintf("%s, %s, %s", data.Address, data.City, data.Region)

	if data.FormattedAddress != "" {
		formattedAddr = data.FormattedAddress
	}

	_, err := repo.AddAddress(s.ctx, repository.AddAddressParams{
		CompanyID:        sql.NullString{String: company.ID, Valid: true},
		Msisdn:           sql.NullString{String: company.Msisdn, Valid: true},
		IsDefault:        true,
		City:             data.City,
		Street:           data.Address,
		Region:           data.Region,
		PostalCode:       sql.NullString{String: data.PostalCode, Valid: true},
		FormattedAddress: sql.NullString{String: formattedAddr, Valid: true},
		Latitude:         sql.NullFloat64{Float64: data.Latitude, Valid: false},
		Longitude:        sql.NullFloat64{Float64: data.Longitude, Valid: false},
	})

	if err != nil {
		slog.Error("AddBillingAddress - AddAddress", "error", err)

		return errors.New("failed to add billing address. please try again")
	}

	return nil
}

func (s *service) SetupAccount(data paymentdto.ReqPaymentDTO) error {
	slog.Info("about to setup payment account", "company", data.CompanyID, "accountNumber", data.PaymentDetails.AccountNumber, "firstName", data.PaymentDetails.FirstName, "lastName", data.PaymentDetails.LastName)

	compservice := userservice.NewCompanyService(s.ctx, s.db)
	company, err := compservice.GetUserCompany()

	if err != nil {
		return err
	}

	repo := s.db.NewRepository()
	errChan := make(chan error)

	defer close(errChan)

	go func() {
		_, err := repo.CreatePaymentAccount(s.ctx, repository.CreatePaymentAccountParams{
			CompanyID:     company.ID,
			AccountNumber: data.PaymentDetails.AccountNumber,
			BankType:      data.PaymentDetails.BankType,
			FirstName:     data.PaymentDetails.FirstName,
			LastName:      data.PaymentDetails.LastName,
			BankCode:      data.PaymentDetails.BankCode,
			BankName:      data.PaymentDetails.BankName,
			BankID:        data.PaymentDetails.BankID,
			BankCurrency:  data.PaymentDetails.BankCurrency,
		})

		if err != nil {
			slog.Error("SetupAccount - CreatePaymentAccount", "error", err)

			errChan <- err
			return
		}

		err = s.AddBillingAddress(company, data.Address)

		if err != nil {
			errChan <- err
			return
		}

		// errChan <- nil

	}()

	if errChan != nil {

		return errors.New("failed to setup payment account. please try again")
	}

	// if err != nil {
	// 	slog.Error("SetupAccount - CreatePaymentAccount", "error", err)

	// 	return errors.New("failed to setup payment account. please try again")
	// }'

	go func() {
		wallet, err := s.setupThirdPartyWallet(company, data)

		if err != nil {
			errChan <- err
			return
		}

		err = s.setUpPayoutAccont(wallet)

		if err != nil {
			errChan <- err
			return
		}

		// errChan <- nil

	}()

	if errChan != nil {

		return errors.New("failed to setup payment account. please try again")
	}

	return nil
}
