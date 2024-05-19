package usercontroller

import (
	"encoding/json"
	"log"
	"net/http"
	userdto "quups-backend/internal/services/user-service/dto"
	userservice "quups-backend/internal/services/user-service/service"
	apiutils "quups-backend/internal/utils/api"
)

// POST: /companies/create
func (c *UserController) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var body *userdto.CreateCompanyParams
	uservice := userservice.New(r.Context(), c.repo)
	response := apiutils.New(w, r)

	err := json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	if err != nil {
		log.Printf("error decoding create company request body")

		res, _ := response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    err.Error(),
		})
		_, _ = w.Write(res)
		return
	}

	newc, err := uservice.CreateCompany(body)

	if err != nil {

		res, _ := response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusBadRequest,
			Results:    nil,
			Message:    err.Error(),
		})
		_, _ = w.Write(res)
		return
	}

	res, _ := response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusCreated,
		Results:    &newc,
		Message:    "success",
	})
	_, _ = w.Write(res)

}
