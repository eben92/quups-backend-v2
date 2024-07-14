package paymentcontroller

import (
	"log"
	"net/http"
	paymentservice "quups-backend/internal/services/payment-service/service"
	apiutils "quups-backend/internal/utils/api"
)

func (c *controller) GetBankList(w http.ResponseWriter, r *http.Request) {

	response := apiutils.New(w, r)
	s := paymentservice.NewPaymentService(r.Context(), c.db)

	res, err := s.GetBankList()

	if err != nil {

		log.Println(err)

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
