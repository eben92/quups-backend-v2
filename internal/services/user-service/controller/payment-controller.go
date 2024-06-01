package usercontroller

import (
	"log"
	"net/http"
	userservice "quups-backend/internal/services/user-service/service"
	apiutils "quups-backend/internal/utils/api"
)

func (c *controller) GetBankList(w http.ResponseWriter, r *http.Request) {

	response := apiutils.New(w, r)

	s := userservice.New(r.Context(), c.db).NewPaymentService()

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
