package paymentservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	paymentdto "quups-backend/internal/services/payment-service/dto"
	"quups-backend/internal/services/payment-service/models"
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
