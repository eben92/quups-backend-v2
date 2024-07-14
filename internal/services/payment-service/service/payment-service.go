package paymentservice

import (
	"encoding/json"
	"fmt"
	"os"
	paymentdto "quups-backend/internal/services/payment-service/dto"
	local_http "quups-backend/internal/utils/http"
)

var (
	PAYSTACK_URL string = os.Getenv("PAYSTACK_BASE_URL")
	ACC_TOKEN    string = os.Getenv("PAYSTACK_SECRET")
)

type apiResponse struct {
	Status  bool
	Message string
	Data    []paymentdto.Bank
}

func (s *service) GetBankList() ([]paymentdto.Bank, error) {

	url := fmt.Sprintf("%s/bank?country=ghana&type=mobile_money", PAYSTACK_URL)
	bearer := fmt.Sprintf("Bearer %s", ACC_TOKEN)

	res, err := local_http.Fetch(url, &local_http.Options{
		Headers: &[][2]string{
			{"Authorization", bearer},
		},
	})

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var response apiResponse

	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		return nil, err
	}

	return response.Data, nil

}
