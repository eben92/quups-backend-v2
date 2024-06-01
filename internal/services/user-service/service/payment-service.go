package userservice

import (
	"encoding/json"
	"fmt"
	"os"
	local_http "quups-backend/internal/utils/http"
	"time"
)

var (
	PAYSTACK_URL string = os.Getenv("PAYSTACK_BASE_URL")
	ACC_TOKEN    string = os.Getenv("PAYSTACK_SECRET")
)

type apiResponse struct {
	Status  bool
	Message string
	Data    *[]Bank
}

type Bank struct {
	Active           bool      `json:"active"`
	Code             string    `json:"code"`
	Country          string    `json:"country"`
	CreatedAt        time.Time `json:"createdAt"`
	Currency         string    `json:"currency"`
	Gateway          *string   `json:"gateway"`
	ID               int       `json:"id"`
	IsDeleted        bool      `json:"is_deleted"`
	Longcode         string    `json:"longcode"`
	Name             string    `json:"name"`
	PayWithBank      bool      `json:"pay_with_bank"`
	Slug             string    `json:"slug"`
	SupportsTransfer bool      `json:"supports_transfer"`
	Type             string    `json:"type"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

func (s *service) GetBankList() (*[]Bank, error) {

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
