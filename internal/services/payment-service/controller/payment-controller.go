package paymentcontroller

import (
	"encoding/json"
	"net/http"
	paymentdto "quups-backend/internal/services/payment-service/dto"
	"quups-backend/internal/services/payment-service/models"
	paymentservice "quups-backend/internal/services/payment-service/service"
	apiutils "quups-backend/internal/utils/api"
)

func (c *controller) GetBankList(w http.ResponseWriter, r *http.Request) {

	bankType := models.BankType(r.URL.Query().Get("type"))

	if bankType == "" {
		bankType = models.MOBILE_MONEY
	}

	response := apiutils.New(w, r)
	s := paymentservice.NewPaymentService(r.Context(), c.db)

	res, err := s.GetBankList(bankType)

	if err != nil {

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Message:    "",
			Results:    nil,
		})

		return
	}

	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success",
		Results:    res,
	})
}

// /resolve-account
func (c *controller) ResolveBankAccount(w http.ResponseWriter, r *http.Request) {

	bankCode := r.URL.Query().Get("bank_code")
	accountNumber := r.URL.Query().Get("account_number")

	if bankCode == "" || accountNumber == "" {
		response := apiutils.New(w, r)

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Message:    "bank_code and account_number are required",
			Results:    nil,
		})

		return
	}

	response := apiutils.New(w, r)
	s := paymentservice.NewPaymentService(r.Context(), c.db)

	res, err := s.ResolveBankAccount(bankCode, accountNumber)

	if err != nil {

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Results:    nil,
		})

		return
	}

	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Message:    "account number resolved successfully",
		Results:    res,
	})
}

func (c *controller) SetupCompanyAccount(w http.ResponseWriter, r *http.Request) {
	var reqBody paymentdto.ReqPaymentDTO

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	response := apiutils.New(w, r)

	if err != nil {
		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid request body",
			Results:    nil,
		})

		return
	}

	defer r.Body.Close()

	err = paymentdto.ValidateReqPaymentDTO(reqBody)

	if err != nil {

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Results:    nil,
		})

		return
	}

	s := paymentservice.NewPaymentService(r.Context(), c.db)

	err = s.SetupAccount(reqBody)

	if err != nil {

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Results:    nil,
		})

		return
	}

	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Message:    "successfully setup payment account ",
		Results:    nil,
	})
}
